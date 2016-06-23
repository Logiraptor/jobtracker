package main

import (
	"jobtracker/app"
	"log"
	"os"
	"strconv"
)

type StdLogger struct{}

func (s StdLogger) Log(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	log.Println(app.Start(app.Context{
		Port:    port,
		AppRoot: ".",
		Logger:  StdLogger{},
	}))
}
