package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listRefundCmd(p *commander.Command) *commander.Command {
	lr := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieves refunds for the provided API token, or payment token",
			Example:   "mollie refunds list --payment=tr_test",
			Execute:   listRefundsAction,
		},
		commander.NoCols(),
	)

	AddPaymentFlag(lr)
	AddFromFlag(lr)
	AddLimitFlag(lr)
	AddEmbedFlag(lr)

	return lr
}

func listRefundsAction(cmd *cobra.Command, args []string) {
	var opts mollie.ListRefundOptions
	{
		opts.Embed = mollie.EmbedValue(ParseStringFromFlags(cmd, EmbedArg))
		opts.Limit = ParseIntFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
	}

	payment := ParseStringFromFlags(cmd, PaymentArg)

	refunds, err := getRefundList(&opts, payment)
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
