package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-bookings/external"
	"hotel-bookings/helpers"
	"hotel-bookings/internal/api"
	"hotel-bookings/internal/interfaces"
	"hotel-bookings/internal/repositories"
	"hotel-bookings/internal/services"
	"os"
)

func ServeHTTP() {
	d := DependencyInjection()
	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	bookingV1 := e.Group("/api/booking/v1")
	bookingV1.POST("/bookings", d.BookingAPI.AddBooking, d.MiddlewareAuthorization)
	bookingV1.GET("/bookings/:id", d.BookingAPI.GetBookingById, d.MiddlewareAuthorization)
	bookingV1.GET("/bookings/user/:id", d.BookingAPI.GetBookingByUserId, d.MiddlewareAuthorization)
	bookingV1.GET("/bookings", d.BookingAPI.GetBookings, d.MiddlewareAdminAuthorization)
	bookingV1.PUT("/bookings/:id", d.BookingAPI.EditBooking, d.MiddlewareAdminAuthorization)
	bookingV1.PATCH("/bookings/:id/status", d.BookingAPI.UpdateBookingStatus)
	bookingV1.GET("/total-booked", d.BookingAPI.GetTotalBookings)

	err := e.Start(":" + os.Getenv("BOOKING_APP_PORT"))
	if err != nil {
		helpers.Logger.Error("failed to start server: ", err)
		return
	}

}

type Dependencies struct {
	External interfaces.IExternal

	BookingAPI interfaces.IBookingAPI
}

func DependencyInjection() *Dependencies {
	ext := external.NewExternal()

	bookingRepo := repositories.NewBookingRepository(helpers.DB)
	bookingSvc := services.NewBookingService(bookingRepo, ext)
	bookingApi := api.NewBookingAPI(bookingSvc)

	return &Dependencies{
		External: ext,

		BookingAPI: bookingApi,
	}
}
