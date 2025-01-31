package external

import (
	bytes3 "bytes"
	"context"
	"encoding/json"
	"fmt"
	"hotel-payments/helpers"
	"net/http"
	"os"
)

func (ex *External) UpdateBookingStatus(ctx context.Context, bookingID, status string) error {
	url := fmt.Sprintf(os.Getenv("BOOKING_URL_UPDATE_STATUS"), bookingID)
	reqBody := map[string]string{"status": status}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	signature := helpers.GenerateSignature(string(jsonBody), os.Getenv("BOOKING_SECRET_KEY"))
	req, err := http.NewRequest(http.MethodPatch, url, bytes3.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Signature", signature)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("resp status code: ", resp.StatusCode)

	return nil
}
