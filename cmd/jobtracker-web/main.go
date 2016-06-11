package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Welcome to Job Tracker")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("App started on port: %s", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
