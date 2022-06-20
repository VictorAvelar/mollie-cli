package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getCustomersCmd(p *commander.Command) *commander.Command {
	gc := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single customer by its ID.",
			Example:   "mollie customers get --id=cs_token",
			Execute:   getCustomerAction,
			PostHook:  printJsonAction,
		},
		customersCols(),
	)

	AddIDFlag(gc, true)

	return gc
}

func getCustomerAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, c, err := app.API.Customers.Get(context.Background(), id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Customers, c, res)

	if verbose {
		app.Logger.Infof("request target: %s", c.Links.Self.Href)
		app.Logger.Infof("request docs: %s", c.Links.Documentation.Href)
	}

	disp := &displayers.MollieCustomer{
		Customer: c,
	}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Customers),
			customersCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
