package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
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
	var opts mollie.ListCustomersOptions
	{
		opts.Limit = ParseIntFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
	}

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	cl, err := API.Customers.List(&opts)
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
