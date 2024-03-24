package main

import (
	"log"

	"github.com/achintya-7/go-template-server/config"
	"github.com/achintya-7/go-template-server/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on port %s", config.Port)

	server := app.NewServer()
	err = server.Start(config.Port)
	if err != nil {
		log.Fatal(err)
	}

}
