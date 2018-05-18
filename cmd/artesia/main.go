package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kennydo/artesia/cmd/artesia/app"
)

var configFilePath = flag.String("config", "configs/config.toml", "Path to the config file")

func main() {
	flag.Parse()

	config := app.NewDefaultConfig()
	if _, err := toml.DecodeFile(*configFilePath, config); err != nil {
		log.Fatal(err)
	}

	s, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
