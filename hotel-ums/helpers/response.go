package helpers

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(e echo.Context, statusCode int, message string, data interface{}) error {
	response := &Response{
		Message: message,
		Data:    data,
	}

	return e.JSON(statusCode, response)
}
