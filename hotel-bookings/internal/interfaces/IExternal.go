package interfaces

import (
	"context"
	"hotel-bookings/external"
)

type IExternal interface {
	ValidateUser(ctx context.Context, token string) (*external.User, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error
}
