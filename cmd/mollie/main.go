package main

import (
	"log"

	"github.com/VictorAvelar/mollie-cli/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal(err)
	}
}
