package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func methods() *commander.Command {
	m := commander.Builder(nil, commander.Config{
		Namespace:          "methods",
		Aliases:            []string{"vendors", "meths"},
		ShortDesc:          "All payment methods that Mollie offers and can be activated",
		PostHook:           printJson,
		PersistentPostHook: printCurl,
	}, getMethodsCols())

	// Add namespace persistent flags.
	AddIncludeFlag(m, true)
	AddPrompterFlag(m, true)

	// Add child commands.
	listPaymentMethodsCmd(m)
	allPaymentMethodsCmd(m)
	getPaymentMethodCmd(m)

	return m
}

func getMethodsCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.methods.all")

	if verbose {
		app.Logger.Info("parsed fields %v", cols)
	}

	return cols
}
