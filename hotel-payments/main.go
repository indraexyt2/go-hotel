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

	// serve http
	go cmd.ServeHTTP()

	// setup kafka
	cmd.ServeKafkaConsumer()

}
