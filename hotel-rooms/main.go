package main

import (
	"hotel-rooms/cmd"
	"hotel-rooms/helpers"
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
