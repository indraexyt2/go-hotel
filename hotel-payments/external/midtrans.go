package external

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TransactionDataResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	Bank              string `json:"bank,omitempty"`
}

func (ex *External) GetMidtransTransactionData(ctx context.Context, orderID string) (*TransactionDataResponse, error) {
	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", orderID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	key := os.Getenv("MIDTRANS_SERVER_KEY")
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(key+":"))
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", auth)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	var transactionDataResponse TransactionDataResponse
	err = json.Unmarshal(body, &transactionDataResponse)
	if err != nil {
		return nil, err
	}

	return &transactionDataResponse, nil
}
