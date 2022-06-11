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

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	_, cl, err := API.Customers.List(context.Background(), &opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", cl.Links.Self.Href)
		logger.Infof("request docs: %s", cl.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomerList{
		CustomersList: cl,
	}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			customersCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
