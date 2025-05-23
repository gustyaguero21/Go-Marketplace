package main

import (
	"go-marketplace/cmd/config"
	"go-marketplace/internal/infrastructure/http/server"
	"log"
)

func main() {

	if envErr := config.LoadEnv(); envErr != nil {
		log.Fatal(envErr)
	}

	server := server.StartServer()

	log.Printf("SERVER STARTED ON PORT%s", config.Port)

	if serverErr := server.Run(config.Port); serverErr != nil {
		log.Fatal("error starting server. Error: ", serverErr)
	}
}
