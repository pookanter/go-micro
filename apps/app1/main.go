package main

import (
	"app1/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/health", handler.HealthCheck)

	port, err := os.LookupEnv("PORT")
	if !err {
		log.Fatal("PORT not set")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
	fmt.Println("App1 is running on port: ", port)
}
