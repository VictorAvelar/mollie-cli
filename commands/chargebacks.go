package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func chargebacks() *commander.Command {
	cb := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "chargebacks",
			ShortDesc:          "Operations with the Chargebacks API",
			Aliases:            []string{"cb"},
			PersistentPostHook: printCurl,
		},
		getChargebacksCols(),
	)

	getChargebacksCmd(cb)
	listChargebacksCmd(cb)

	return cb
}

func getChargebacksCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.chargebacks.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}
