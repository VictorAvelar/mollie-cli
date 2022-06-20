package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func allRefundsCmd(p *commander.Command) *commander.Command {
	ar := commander.Builder(
		p,
		commander.Config{
			Namespace: "all",
			ShortDesc: "List all refunds for the account",
			Aliases:   []string{"all", "complete"},
			Execute:   allRefundsAction,
			Example:   "mollie refunds all",
		},
		refundsCols(),
	)

	AddEmbedFlag(ar)

	return ar
}

func allRefundsAction(cmd *cobra.Command, args []string) {
	var opts mollie.ListRefundOptions
	{
		opts.Embed = mollie.EmbedValue(ParseStringFromFlags(cmd, EmbedArg))
	}

	refunds, err := getRefundList(&opts, "")
	if err != nil {
		app.Logger.Fatal(err)
	}

	if verbose {
		app.Logger.Infof("retrieved %d refunds", refunds.Count)
		app.Logger.Infof("request target: %s", refunds.Links.Self.Href)
		app.Logger.Infof("request docs: %s", refunds.Links.Documentation.Href)
	}

	disp := &displayers.MollieRefundList{RefundList: refunds}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(parseFieldsFromFlag(cmd, Refunds), refundsCols()),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
