package main

import (
	"hotel-bookings/cmd"
	"hotel-bookings/helpers"
)

func main() {
	// setup logger
	helpers.SetupLogger()

	// setup db
	helpers.SetupDB()

	// setup redis
	helpers.SetupRedis()

	// serve http
	cmd.ServeHTTP()
}
