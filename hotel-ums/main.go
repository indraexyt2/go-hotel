package main

import (
	"hotel-ums/cmd"
	"hotel-ums/helpers"
)

func main() {
	// Setup Logger
	helpers.SetupLogger()

	// Setup DB
	helpers.SetupDB()

	// Setup Redis
	helpers.SetupRedis()

	// Serve HTTP
	cmd.ServeHTTP()
}
