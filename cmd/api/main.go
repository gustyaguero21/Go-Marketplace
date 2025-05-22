package main

import (
	"go-marketplace/cmd/config"
	"go-marketplace/internal/infrastructure/database"
	"go-marketplace/internal/infrastructure/http/server"
	"log"
)

func main() {

	if envErr := config.LoadEnv(); envErr != nil {
		log.Fatal(envErr)
	}

	_, err := database.DatabaseConn()
	if err != nil {
		log.Fatal(err)
	}

	server := server.StartServer()

	log.Printf("SERVER STARTED ON PORT%s", config.Port)

	if serverErr := server.Run(config.Port); serverErr != nil {
		log.Fatal("error starting server. Error: ", serverErr)
	}
}
