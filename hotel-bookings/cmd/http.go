package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-bookings/external"
	"hotel-bookings/internal/interfaces"
	"os"
)

func ServeHTTP() {
	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	err := e.Start(":" + os.Getenv("ROOM_APP_PORT"))
	if err != nil {
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
