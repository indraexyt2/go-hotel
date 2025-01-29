package external

import (
	"context"
	"encoding/json"
	"hotel-rooms/internal/models"
	"net/http"
	"os"
)

type TotalBookedResponse struct {
	Message string                      `json:"message"`
	Data    []models.RoomBookedResponse `json:"data"`
}

func (ex *External) GetTotalBooked(ctx context.Context) ([]models.RoomBookedResponse, error) {
	url := os.Getenv("BOOKING_URL_TOTAL_BOOKED")
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	responseData := &TotalBookedResponse{}
	err = json.NewDecoder(response.Body).Decode(responseData)
	if err != nil {
		return nil, err
	}

	return responseData.Data, nil
}
