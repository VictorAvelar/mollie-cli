package commands

import (
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getPaymentCmd(p *commander.Command) *commander.Command {
	gp := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single payment object by its payment token.",
			Execute:   getPaymentAction,
			Example:   "mollie payments get --id=tr_token",
		},
		getPaymentCols(),
	)

	AddIDFlag(gp, true)

	return gp
}

func getPaymentAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		logger.Infof("retrieving payment with id (token) %s", id)
	}

	p, err := API.Payments.Get(id, nil)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
	}

	if json {
		PrintJsonP(p)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			getPaymentCols(),
		),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
