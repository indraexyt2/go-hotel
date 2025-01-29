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

	roomV1.POST("/rooms", d.RoomsAPI.AddRoom, d.MiddlewareAdminAuthorization)
	roomV1.GET("/rooms", d.RoomsAPI.GetRooms)
	roomV1.GET("/rooms/:id", d.RoomsAPI.GetRoomDetails)
	roomV1.PUT("/rooms/:id", d.RoomsAPI.EditRoom, d.MiddlewareAdminAuthorization)
	roomV1.DELETE("/rooms/:id", d.RoomsAPI.DeleteRoom, d.MiddlewareAdminAuthorization)
	roomV1.GET("/rooms-availability", d.RoomsAPI.GetRoomAvailability)

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
	RoomsAPI        interfaces.IRoomsAPI
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

	roomsRepo := repositories.NewRoomsRepository(helpers.DB, helpers.RedisClient)
	roomsSvc := services.NewRoomService(roomsRepo)
	roomsApi := api.NewRoomsAPI(roomsSvc, ext)

	return &Dependencies{
		External: ext,

		RoomTypesAPI:    roomTypesApi,
		RoomFeaturesAPI: roomFeaturesApi,
		RoomPhotosAPI:   roomPhotosApi,
		RoomsAPI:        roomsApi,
	}
}
