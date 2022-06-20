package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func invoices() *commander.Command {
	i := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "invoices",
			Example:            "mollie invoices",
			ShortDesc:          "Operations over Mollie's Invoices API.",
			PostHook:           printJson,
			PersistentPostHook: printCurl,
		},
		invoicesCols(),
	)

	listInvoicesCmd(i)
	getInvoicesCmd(i)

	return i
}

func invoicesCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.invoices.all")

	if verbose {
		app.Logger.Info("parsed fields %v", cols)
	}

	return cols
}
