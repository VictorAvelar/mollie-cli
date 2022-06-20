package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func captures() *commander.Command {
	c := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "captures",
			ShortDesc:          "Operations with Captures API.",
			PostHook:           printJson,
			PersistentPostHook: printCurl,
		},
		getCapturesCols(),
	)

	listCapturesCmd(c)
	getCapturesCmd(c)

	return c
}

func getCapturesCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.captures.all")

	if verbose {
		app.Logger.Info("parsed fields %v", cols)
	}

	return cols
}
