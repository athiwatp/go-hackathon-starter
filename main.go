package main

import (
	"log"

	"github.com/niranjan92/go_hackathon_starter/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
