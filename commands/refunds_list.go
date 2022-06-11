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
		noCols,
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

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	refunds, err := getRefundList(&opts, payment)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("retrieved %d refunds", refunds.Count)
		logger.Infof("request target: %s", refunds.Links.Self.Href)
		logger.Infof("request docs: %s", refunds.Links.Documentation.Href)
	}

	disp := &displayers.MollieRefundList{RefundList: refunds}

	err = printer.Display(
		disp,
		display.FilterColumns(parseFieldsFromFlag(cmd), refundsCols()),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
