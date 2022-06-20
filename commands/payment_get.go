package commands

import (
	"context"

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
			PostHook:  printJsonAction,
		},
		getPaymentCols(),
	)

	AddIDFlag(gp, true)

	return gp
}

func getPaymentAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, p, err := app.API.Payments.Get(context.Background(), id, nil)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Payments, p, res)

	if verbose {
		app.Logger.Infof("request target: %s", p.Links.Self.Href)
		app.Logger.Infof("request docs: %s", p.Links.Documentation.Href)
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
