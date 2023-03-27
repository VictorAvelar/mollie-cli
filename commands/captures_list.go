package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listCapturesCmd(p *commander.Command) *commander.Command {
	lc := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all captures for a certain payment.",
			LongDesc: `Retrieve all captures for a certain payment.
Captures are used for payments that have the authorize-then-capture flow.
The only payment methods at the moment that have this flow are Klarna Pay
later and Klarna Slice it.`,
			Execute:  listCapturesActions,
			Example:  "mollie captures list --payment tr_example",
			PostHook: printJsonAction,
		},
		getCapturesCols(),
	)

	commander.AddFlag(lc, commander.FlagConfig{
		Name:     PaymentArg,
		Usage:    "the payment id/token",
		Required: true,
	})

	return lc
}

func listCapturesActions(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)

	res, captures, err := app.API.Captures.List(context.Background(), payment)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Captures, captures, res)

	if verbose {
		app.Logger.Infof("request target: %s", captures.Links.Self.Href)
		app.Logger.Infof("request docs: %s", captures.Links.Documentation.Href)
	}

	disp := &displayers.MollieCapturesList{
		CapturesList: captures,
	}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Payments),
			getCapturesCols(),
		),
	)

	if err != nil {
		app.Logger.Fatal(err)
	}
}
