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
	roomV1.GET("/room-types/:id", d.RoomTypesAPI.GetRoomTypesDetails)
	roomV1.POST("/room-types", d.RoomTypesAPI.AddRoomType, d.MiddlewareAdminAuthorization)
	roomV1.PUT("/room-types/:id", d.RoomTypesAPI.UpdateRoomType, d.MiddlewareAdminAuthorization)
	roomV1.DELETE("/room-types/:id", d.RoomTypesAPI.DeleteRoomType, d.MiddlewareAdminAuthorization)

	roomV1.POST("/room-types/features", d.RoomFeaturesAPI.AddRoomFeature, d.MiddlewareAdminAuthorization)
	roomV1.GET("/room-types/:roomTypeId/features", d.RoomFeaturesAPI.GetAllRoomFeatures)
	roomV1.PUT("/room-types/features/:id", d.RoomFeaturesAPI.EditRoomFeature, d.MiddlewareAdminAuthorization)
	roomV1.DELETE("/room-types/features/:id", d.RoomFeaturesAPI.DeleteRoomFeature, d.MiddlewareAdminAuthorization)

	roomV1.POST("/room-types/:roomTypeId/photos", d.RoomPhotosAPI.AddRoomTypePhotos, d.MiddlewareAdminAuthorization)
	roomV1.GET("/room-types/:roomTypeId/photos", d.RoomPhotosAPI.GetRoomTypePhotos)
	roomV1.DELETE("/room-types/photos/:id", d.RoomPhotosAPI.DeleteRoomTypePhoto)

	err := e.Start(":" + os.Getenv("ROOM_APP_PORT"))
	if err != nil {
		return
	}
}

type Dependencies struct {
	External *external.External

	RoomTypesAPI    interfaces.IRoomTypesAPI
	RoomFeaturesAPI interfaces.IRoomFeaturesAPI
	RoomPhotosAPI   interfaces.IRoomPhotosAPI
}

func DependencyInjection() *Dependencies {
	ext := external.NewExternal()

	roomTypesRepo := repositories.NewRoomTypesRepository(helpers.DB)
	roomTypesSvc := services.NewRoomTypesService(roomTypesRepo)
	roomTypesApi := api.NewRoomTypesAPI(roomTypesSvc)

	roomFeaturesRepo := repositories.NewRoomFeaturesRepository(helpers.DB)
	roomFeaturesSvc := services.NewRoomFeaturesService(roomFeaturesRepo)
	roomFeaturesApi := api.NewRoomFeaturesAPI(roomFeaturesSvc)

	roomPhotosRepo := repositories.NewRoomPhotosRepository(helpers.DB)
	roomPhotosSvc := services.NewRoomPhotosService(roomPhotosRepo)
	roomPhotosApi := api.NewRoomPhotosAPI(roomPhotosSvc)

	return &Dependencies{
		External: ext,

		RoomTypesAPI:    roomTypesApi,
		RoomFeaturesAPI: roomFeaturesApi,
		RoomPhotosAPI:   roomPhotosApi,
	}
}
