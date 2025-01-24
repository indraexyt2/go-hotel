package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/api"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/repositories"
	"hotel-ums/internal/services"
	"os"
)

func ServeHTTP() {
	d := DependencyInjection()

	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	umsV1 := e.Group("/api/ums/v1")
	umsV1.POST("/register", d.RegisterAPI.RegisterNewUser)
	umsV1.PUT("/email-verification/:token", d.RegisterAPI.EmailVerification)
	umsV1.GET("/email-verification", d.RegisterAPI.ResendEmailVerification)
	umsV1.POST("/login", d.LoginAPI.Login)

	err := e.Start(":" + os.Getenv("UMS_APP_PORT"))
	if err != nil {
		return
	}
}

type Dependencies struct {
	RegisterAPI interfaces.IUserRegisterAPI
	LoginAPI    interfaces.IUserLoginAPI
}

func DependencyInjection() *Dependencies {
	userRepo := repositories.NewUserRepository(helpers.DB, helpers.RedisClient)

	registerSvc := services.NewRegisterService(userRepo)
	registerApi := api.NewRegisterAPI(registerSvc)

	loginSvc := services.NewLoginService(userRepo)
	loginApi := api.NewLoginAPI(loginSvc)

	return &Dependencies{
		RegisterAPI: registerApi,
		LoginAPI:    loginApi,
	}
}
