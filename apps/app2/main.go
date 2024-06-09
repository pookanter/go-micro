package main

import (
	"go-micro/apps/app1/handler"
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
