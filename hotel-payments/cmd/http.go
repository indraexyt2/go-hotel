package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-payments/external"
	"hotel-payments/helpers"
	"hotel-payments/internal/api"
	"hotel-payments/internal/interfaces"
	"hotel-payments/internal/repositories"
	"hotel-payments/internal/services"
	"os"
)

func ServeHTTP() {
	d := DependencyInjection()
	e := echo.New()
	e.GET("/ping", func(e echo.Context) error {
		return e.String(200, "pong")
	})

	paymentV1 := e.Group("/api/payment/v1")
	paymentV1.POST("/midtrans/notification", d.PaymentsAPI.ProcessPaymentCallback)
	paymentV1.POST("/midtrans/refund", d.PaymentsAPI.RefundPayment, d.MiddlewareAuthorization)

	err := e.Start(":" + os.Getenv("PAYMENT_APP_PORT"))
	if err != nil {
		helpers.Logger.Error("failed to start server: ", err)
		return
	}

}

type Dependencies struct {
	External interfaces.IExternal

	PaymentsAPI interfaces.IPaymentsAPI
}

func DependencyInjection() *Dependencies {
	ext := external.NewExternal()

	paymentRepo := repositories.NewPaymentsRepository(helpers.DB)
	paymentSvc := services.NewPaymentService(paymentRepo, ext)
	paymentApi := api.NewPaymentAPI(paymentSvc, helpers.SnapClient(), helpers.CoreClient(), ext)

	return &Dependencies{
		External: ext,

		PaymentsAPI: paymentApi,
	}
}
