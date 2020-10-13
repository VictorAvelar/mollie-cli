package main

import (
	"github.com/VictorAvelar/mollie-cli/commands"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := commands.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
