package main

import (
	"go-marketplace/cmd/config"
	"go-marketplace/internal/infrastructure/http/server"
	"log"
)

func main() {

	server := server.StartServer()

	log.Printf("SERVER STARTED ON PORT%s", config.Port)

	if err := server.Run(config.Port); err != nil {
		log.Fatal("error starting server. Error: ", err)
	}
}
