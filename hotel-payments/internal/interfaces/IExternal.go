package interfaces

import (
	"context"
	"hotel-payments/external"
)

type IExternal interface {
	ValidateUser(ctx context.Context, token string) (*external.User, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error
	GetMidtransTransactionData(ctx context.Context, orderID string) (*external.TransactionDataResponse, error)
}
