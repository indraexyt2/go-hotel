package api

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type RoomPhotosAPI struct {
	RoomPhotosSVC interfaces.IRoomPhotosService
}

func NewRoomPhotosAPI(roomPhotosSVC interfaces.IRoomPhotosService) *RoomPhotosAPI {
	return &RoomPhotosAPI{RoomPhotosSVC: roomPhotosSVC}
}

func (api *RoomPhotosAPI) AddRoomTypePhotos(e echo.Context) error {
	var (
		log = helpers.Logger
		req []models.RoomPhoto
	)

	roomTypeId := e.Param("roomTypeId")
	roomTypeIdInt, err := strconv.Atoi(roomTypeId)
	if err != nil {
		log.Error("failed to convert room type id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = e.Request().ParseMultipartForm(5 << 20)
	if err != nil {
		log.Error("failed to parse multipart form: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	isPrimaryValues := e.Request().MultipartForm.Value["is_primary"]
	if len(isPrimaryValues) == 0 {
		log.Error("failed to get is primary values: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, "failed to get is primary values", nil)
	}

	photos := e.Request().MultipartForm.File["photos"]

	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	uploadDir := filepath.Join(dir, "../../uploads/photos/")
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			log.Error("failed to create directory: ", err)
			return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
		}
	}

	for _, photoHeader := range photos {
		photo, err := photoHeader.Open()
		if err != nil {
			log.Error("failed to open photo: ", err)
			return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
		}
		defer photo.Close()

		filePath := filepath.Join(uploadDir, photoHeader.Filename)
		relativePath := filepath.Join("/uploads/photos/", photoHeader.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			log.Error("failed to create file: ", err)
			return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, photo); err != nil {
			log.Error("failed to copy photo: ", err)
			return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
		}

		isPrimary := false
		for _, primaryFile := range isPrimaryValues {
			if primaryFile == photoHeader.Filename {
				isPrimary = true
				break
			}
		}

		photoData := &models.RoomPhoto{
			RoomTypeID: uint(roomTypeIdInt),
			FilePath:   relativePath,
			IsPrimary:  isPrimary,
		}

		req = append(req, *photoData)
	}

	err = api.RoomPhotosSVC.AddRoomTypePhotos(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to add room type photos: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomPhotosAPI) GetRoomTypePhotos(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	roomTypeId := e.Param("roomTypeId")
	roomTypeIdInt, err := strconv.Atoi(roomTypeId)
	if err != nil {
		log.Error("failed to convert room type id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.RoomPhotosSVC.GetRoomTypePhotos(e.Request().Context(), roomTypeIdInt)
	if err != nil {
		log.Error("failed to get room type photos: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomPhotosAPI) DeleteRoomTypePhoto(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = api.RoomPhotosSVC.DeletePhotos(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to delete room type photo: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}
