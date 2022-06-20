package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getChargebacksCmd(p *commander.Command) *commander.Command {
	gcb := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.",
			LongDesc: `Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.
A debit to a depositor's account for an item that has been previously credited, as for a returned bad check.`,
			Execute:  getChargebackAction,
			Example:  "mollie chargebacks get --id=cb_token --embed=payments",
			PostHook: printJsonAction,
		},
		getChargebacksCols(),
	)
	commander.AddFlag(gcb, commander.FlagConfig{
		Name:     PaymentArg,
		Usage:    "original payment id/token",
		Required: true,
	})
	commander.AddFlag(gcb, commander.FlagConfig{
		Name:     IDArg,
		Usage:    "the chargeback id/token",
		Required: true,
	})
	commander.AddFlag(gcb, commander.FlagConfig{
		Name:  EmbedArg,
		Usage: "a comma separated list of embedded resources",
	})

	return gcb
}

func getChargebackAction(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	chargeback := ParseStringFromFlags(cmd, IDArg)
	embed := ParseStringFromFlags(cmd, EmbedArg)

	var opts mollie.ChargebackOptions
	if len(embed) > 0 {
		opts = mollie.ChargebackOptions{
			Embed: embed,
		}
	}

	res, cb, err := app.API.Chargebacks.Get(context.Background(), payment, chargeback, &opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Chargebacks, cb, res)

	if verbose {
		app.Logger.Infof("request target: %s", cb.Links.Self.Href)
		app.Logger.Infof("request docs: %s", cb.Links.Documentation.Href)
	}

	disp := &displayers.MollieChargeback{Chargeback: cb}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Chargebacks),
			getChargebacksCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
