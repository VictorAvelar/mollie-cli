package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listChargebacksCmd(p *commander.Command) *commander.Command {
	lcb := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all received chargebacks",
			LongDesc: `Retrieve all received chargebacks. If the payment-specific endpoint is used, only chargebacks
for that specific payment are returned.`,
			Execute:  listChargebackAction,
			Example:  "mollie chargebacks list --embed=payments",
			PostHook: printJsonAction,
		},
		getChargebacksCols(),
	)
	commander.AddFlag(lcb, commander.FlagConfig{
		Name:  EmbedArg,
		Usage: "a comma separated list of embedded resources",
	})

	return lcb
}

func listChargebackAction(cmd *cobra.Command, args []string) {
	embed := ParseStringFromFlags(cmd, EmbedArg)

	var opt mollie.ChargebacksListOptions
	{
		opt.Embed = embed
	}

	res, cbs, err := app.API.Chargebacks.List(context.Background(), &opt)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Chargebacks, cbs, res)

	if verbose {
		app.Logger.Infof("response with %d chargebacks", cbs.Count)
		app.Logger.Infof("request target: %s", cbs.Links.Self.Href)
		app.Logger.Infof("request docs: %s", cbs.Links.Documentation.Href)
	}

	disp := displayers.MollieChargebackList{ChargebacksList: cbs}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Chargebacks),
			getChargebacksCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
