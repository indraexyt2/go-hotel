package main

import (
	"hotel-payments/cmd"
	"hotel-payments/helpers"
)

func main() {
	// setup logger
	helpers.SetupLogger()

	// setup db
	helpers.SetupDB()

	// setup redis
	helpers.SetupRedis()

	// setup midtrans
	helpers.SetupMidtrans()

	// setup kafka
	go cmd.ServeKafkaConsumer()

	// serve http
	cmd.ServeHTTP()
}
