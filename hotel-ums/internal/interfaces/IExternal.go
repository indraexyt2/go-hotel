package interfaces

import "context"

type IExternal interface {
	SendMessageNotification(ctx context.Context, message []byte) error
}
