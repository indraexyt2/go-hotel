package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type ProfileAPI struct {
	ProfileService interfaces.IProfileService
	UserRepository interfaces.IUserRepository
}

func NewProfileAPI(profSvc interfaces.IProfileService, userRepo interfaces.IUserRepository) *ProfileAPI {
	return &ProfileAPI{
		ProfileService: profSvc,
		UserRepository: userRepo,
	}
}

func (api *ProfileAPI) UpdateUserProfile(e echo.Context) error {
	var (
		log    = helpers.Logger
		req    = &models.User{}
		userID int
	)

	token := e.Get("token")
	switch token.(type) {
	case *helpers.Claims:
		claimsToken, ok := token.(*helpers.Claims)
		if !ok {
			log.Error("error getting token")
			return helpers.SendResponse(e, 500, "server error", nil)
		}
		userID = claimsToken.UserID
	case *models.GoogleUserInfo:
		claimsToken, ok := token.(*models.GoogleUserInfo)
		if !ok {
			log.Error("error getting token")
			return helpers.SendResponse(e, 500, "server error", nil)
		}

		user, err := api.UserRepository.GetUserByEmail(e.Request().Context(), claimsToken.Email)
		if err != nil {
			log.Error("failed to get user: ", err)
			return helpers.SendResponse(e, 500, err.Error(), nil)
		}
		userID = user.ID
	default:
		log.Error("error getting token")
		return helpers.SendResponse(e, 500, "server error", nil)
	}

	photoSrc, err := e.FormFile("file")
	if err != nil {
		log.Error("failed to get photo: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	uploadPath := filepath.Join(dir, "../../uploads/profile")
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			log.Error("failed to create directory: ", err)
			return helpers.SendResponse(e, 500, "failed to create directory", nil)
		}
	}

	filePath := fmt.Sprintf("%s/profile-%s-%d", uploadPath, photoSrc.Filename, time.Now().Unix())
	relativePath := fmt.Sprintf("/uploads/profile/%s", photoSrc.Filename)
	src, err := photoSrc.Open()
	if err != nil {
		log.Error("failed to open file: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		log.Error("failed to create file: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Error("failed to copy file: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("Failed to bind user: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	resp, err := api.ProfileService.UpdateUserProfile(e.Request().Context(), req, relativePath, userID)
	if err != nil {
		log.Error("Failed to update user profile: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)

}
