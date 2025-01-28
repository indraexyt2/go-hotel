package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/external"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/api"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/repositories"
	"hotel-rooms/internal/services"
	"os"
)

func ServeHTTP() {
	d := DependencyInjection()
	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	roomV1 := e.Group("/api/room/v1")
	roomV1.GET("/room-types", d.RoomTypesAPI.GetAllRoomTypes)

	err := e.Start(":" + os.Getenv("ROOM_APP_PORT"))
	if err != nil {
		return
	}
}

type Dependencies struct {
	External *external.External

	RoomTypesAPI interfaces.IRoomTypesAPI
}

func DependencyInjection() *Dependencies {
	ext := external.NewExternal()

	roomTypesRepo := repositories.NewRoomTypesRepository(helpers.DB)
	roomTypesSvc := services.NewRoomTypesService(roomTypesRepo)
	roomTypesApi := api.NewRoomTypesAPI(roomTypesSvc)

	return &Dependencies{
		External: ext,

		RoomTypesAPI: roomTypesApi,
	}
}
