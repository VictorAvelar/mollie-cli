package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listCustomerCmd(p *commander.Command) *commander.Command {
	lc := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieves all customers created.",
			Example:   "mollie customers list",
			Execute:   listCustomerActions,
			PostHook:  printJsonAction,
		},
		customersCols(),
	)

	AddLimitFlag(lc)
	AddFromFlag(lc)

	return lc
}

func listCustomerActions(cmd *cobra.Command, args []string) {
	var opts mollie.CustomersListOptions
	{
		opts.Limit = ParseIntFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
	}

	res, cl, err := app.API.Customers.List(context.Background(), &opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Customers, cl, res)

	if verbose {
		app.Logger.Infof("request target: %s", cl.Links.Self.Href)
		app.Logger.Infof("request docs: %s", cl.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomerList{
		CustomersList: cl,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Customers),
			customersCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
