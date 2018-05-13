package main

import (
	"github.com/kennydo/artesia/cmd/artesia/app"
)

func main() {
	s, _ := app.NewServer()
	s.Run()
}
