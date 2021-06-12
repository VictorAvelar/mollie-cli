package commands

import (
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func cancelPaymentCmd(p *commander.Command) *commander.Command {
	cp := commander.Builder(
		p,
		commander.Config{
			Namespace: "cancel",
			ShortDesc: "Cancel a payment by its payment token.",
			Execute:   cancelPaymentAction,
			Example:   "mollie payments cancel --id=tr_token",
		},
		getPaymentCols(),
	)

	AddIDFlag(cp, true)

	return cp
}

func cancelPaymentAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		logger.Infof("canceling payment with id (token) %s", id)
	}

	p, err := API.Payments.Cancel(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
		logger.Infof("payment successfully cancelled")
		logger.Infof("cancellation processed at %s", p.CanceledAt)
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
