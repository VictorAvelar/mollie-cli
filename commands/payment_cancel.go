package commands

import (
	"context"

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
			PostHook:  printJsonAction,
		},
		getPaymentCols(),
	)

	AddIDFlag(cp, true)

	return cp
}

func cancelPaymentAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, p, err := app.API.Payments.Cancel(context.Background(), id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Payments, id, res)

	if verbose {
		app.Logger.Infof("request target: %s", p.Links.Self.Href)
		app.Logger.Infof("request docs: %s", p.Links.Documentation.Href)
		app.Logger.Infof("payment successfully cancelled")
		app.Logger.Infof("cancellation processed at %s", p.CanceledAt)
	}

	disp := displayers.MolliePayment{Payment: p}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Payments),
			getPaymentCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
