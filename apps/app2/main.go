package main

import (
	"app2/handler"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/health", handler.HealthCheck)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
