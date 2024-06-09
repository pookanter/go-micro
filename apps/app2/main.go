package main

import (
	"app2/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.HealthCheck)

	port, err := os.LookupEnv("PORT")
	if !err {
		log.Fatal("PORT not set")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("App2 is running on port: ", port)

	log.Fatal(server.ListenAndServe())
}
