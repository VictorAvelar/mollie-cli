package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func customers() *commander.Command {
	c := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "customers",
			ShortDesc:          "Operations with customers API.",
			Aliases:            []string{"cust"},
			PostHook:           printJsonAction,
			PersistentPostHook: printCurl,
		},
		customersCols(),
	)

	getCustomersCmd(c)
	listCustomerCmd(c)
	createCustomerCmd(c)
	updateCustomerCmd(c)
	deleteCustomerCmd(c)

	return c
}

func customersCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.customers.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}
