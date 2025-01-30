package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hotel-rooms/internal/models"
	"net/http"
	"os"
)

type TotalBookedResponse struct {
	Message string                      `json:"message"`
	Data    []models.RoomBookedResponse `json:"data"`
}

func (ex *External) GetTotalBooked(ctx context.Context, checkinDate string, checkoutDate string) ([]models.RoomBookedResponse, error) {
	url := os.Getenv("BOOKING_URL_TOTAL_BOOKED") + "?checkinDate=" + checkinDate + "&checkoutDate=" + checkoutDate
	fmt.Println("url: ", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed to get total booked: %s", response.Status))
	}

	fmt.Println("response body: ", response.Body)

	responseData := &TotalBookedResponse{}
	err = json.NewDecoder(response.Body).Decode(responseData)
	if err != nil {
		return nil, err
	}

	return responseData.Data, nil
}
