package main

import (
	"github.com/VictorAvelar/mollie-cli/commands"
	"github.com/sirupsen/logrus"
)

func main() {
	var logger *logrus.Logger
	{
		logrus.SetReportCaller(true)
		logger = logrus.New()
	}

	if err := commands.Execute(); err != nil {
		logger.Fatal(err)
	}
}
