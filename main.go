package main

import (
	"log"

	"github.com/unedtamps/go-backend/config"
)

func main() {
	server, err := config.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
