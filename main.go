package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	app "github.com/whoismath/ademir/app"
	config "github.com/whoismath/ademir/config"
)

func main() {

	c, err := config.Setup()

	if err != nil {
		log.Fatal(err)
	}

	t := make(chan string)
	defer close(t)

	a := app.CreateApp(c, t)
	a.Start()
}
