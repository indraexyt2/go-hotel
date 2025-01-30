package helpers

import (
	midtrans2 "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
)

func SnapClient() *snap.Client {
	var s *snap.Client
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans2.Sandbox)
	return s
}
