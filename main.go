package main

import (
	"log"

	"github.com/whoismath/ademir/app"
	"github.com/whoismath/ademir/config"
)

func main() {
	c, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	a := app.CreateApp(c)
	a.Start()
}
