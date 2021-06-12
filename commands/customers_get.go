package commands

import (
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
		},
		customersCols(),
	)

	AddIDFlag(gc, true)

	return gc
}

func getCustomerAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
	}

	c, err := API.Customers.Get(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", c.Links.Self.Href)
		logger.Infof("request docs: %s", c.Links.Documentation.Href)
	}

	disp := &displayers.MollieCustomer{
		Customer: c,
	}

	err = printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			customersCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
