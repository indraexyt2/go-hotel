package helpers

import (
	midtrans2 "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/veritrans/go-midtrans"
	"os"
)

func SnapClient() *snap.Client {
	var env midtrans2.EnvironmentType
	switch os.Getenv("ENV_MODE") {
	case "production":
		env = midtrans2.Production
	case "sandbox":
		env = midtrans2.Sandbox
	}

	s := &snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), env)
	return s
}

func CoreClient() *midtrans.CoreGateway {
	var env midtrans.EnvironmentType
	switch os.Getenv("ENV_MODE") {
	case "production":
		env = midtrans.Production
	case "sandbox":
		env = midtrans.Sandbox
	}

	midClient := midtrans.NewClient()
	midClient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midClient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midClient.APIEnvType = env

	return &midtrans.CoreGateway{
		Client: midClient,
	}
}
