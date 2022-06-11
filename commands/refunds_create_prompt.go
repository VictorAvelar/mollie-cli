package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func promptCreateRefundCmd(p *commander.Command) *commander.Command {
	return commander.Builder(
		p,
		commander.Config{
			Namespace: "prompt",
			ShortDesc: "prompts the user the information necessary to create a refund.",
			Example:   "mollie refunds create prompt",
			Execute:   promptCreateRefundAction,
		},
		refundsCols(),
	)
}

func promptCreateRefundAction(cmd *cobra.Command, args []string) {
	payment := promptStringClean("payment id", "")
	if payment == "" {
		logger.Fatal("a payment id is required to create a refund")
	}

	var r mollie.Refund
	{
		r = mollie.Refund{
			Amount:      promptAmount(),
			Description: promptStringClean("refund description", ""),
			Metadata:    promptStringClean("additional metadata", ""),
		}
	}

	_, rs, err := API.Refunds.Create(context.Background(), payment, r, nil)
	if err != nil {
		logger.Errorf("%+v", rs)
		logger.Errorf("%+v", r)
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("refund for payment %s created", payment)
		logger.Infof("request target: %s", rs.Links.Self.Href)
		logger.Infof("request docs: %s", rs.Links.Documentation.Href)
	}

	if json {
		printJSONP(rs)
	}

	disp := displayers.MollieRefund{Refund: rs}

	err = printer.Display(
		&disp,
		display.FilterColumns(parseFieldsFromFlag(cmd), refundsCols()),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
