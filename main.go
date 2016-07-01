package main

import (
	"jobtracker/app"

	log "github.com/Sirupsen/logrus"

	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	logger := log.New()
	logger.Level = log.DebugLevel
	log.Println(app.Start(app.Context{
		Port:    port,
		AppRoot: ".",
		Logger:  logger,
	}))
}
