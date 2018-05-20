package main

import (
	"log"

	"github.com/kennydo/artesia/cmd/artesia/app"
)

func main() {
	_, err := app.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
