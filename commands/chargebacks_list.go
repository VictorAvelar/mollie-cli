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
			Execute: listChargebackAction,
			Example: "mollie chargebacks list --embed=payments",
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

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	var opt mollie.ChargebacksListOptions
	{
		opt.Embed = embed
	}

	_, cbs, err := API.Chargebacks.List(context.Background(), &opt)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("response with %d chargebacks", cbs.Count)
		logger.Infof("request target: %s", cbs.Links.Self.Href)
		logger.Infof("request docs: %s", cbs.Links.Documentation.Href)
	}

	disp := displayers.MollieChargebackList{ChargebacksList: cbs}

	err = printer.Display(
		&disp,
		display.FilterColumns(parseFieldsFromFlag(cmd), getChargebacksCols()),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
