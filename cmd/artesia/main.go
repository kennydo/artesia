package main

import (
	"log"

	"github.com/kennydo/artesia/cmd/artesia/app"
)

func main() {
	config, err := app.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := app.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
