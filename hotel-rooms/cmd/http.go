package cmd

import (
	"github.com/labstack/echo/v4"
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
}

func DependencyInjection() *Dependencies {
	return &Dependencies{}
}
