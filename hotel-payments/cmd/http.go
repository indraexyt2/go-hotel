package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-payments/external"
	"hotel-payments/helpers"
	"hotel-payments/internal/interfaces"
	"os"
)

func ServeHTTP() {
	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	err := e.Start(":" + os.Getenv("PAYMENT_APP_PORT"))
	if err != nil {
		helpers.Logger.Error("failed to start server: ", err)
		return
	}

}

type Dependencies struct {
	External interfaces.IExternal
}

func DependencyInjection() *Dependencies {
	ext := external.NewExternal()

	return &Dependencies{
		External: ext,
	}
}
