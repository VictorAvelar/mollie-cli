package main

import (
	"log"

	"github.com/VictorAvelar/mollie-cli/commands"
)

func main() {
	app := commands.New()
	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
