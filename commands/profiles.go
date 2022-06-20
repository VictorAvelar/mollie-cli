package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func profile() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "profiles",
			ShortDesc:          "In order to process payments, you need to create a website profile",
			PostHook:           printJson,
			PersistentPostHook: printCurl,
		},
		getProfileCols(),
	)

	getProfileCmd(p)
	currentProfileCmd(p)

	return p
}

func getProfileCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.profiles.all")

	if verbose {
		app.Logger.Info("parsed fields %v", cols)
	}

	return cols
}
