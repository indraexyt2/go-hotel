package main

import (
	"hotel-notification/cmd"
	"hotel-notification/helpers"
)

func main() {
	// setup logger
	helpers.SetupLogger()

	// setup db
	helpers.SetupDB()

	// setup redis
	helpers.SetupRedis()

	// serve kafka
	cmd.ServeKafkaConsumer()
}
