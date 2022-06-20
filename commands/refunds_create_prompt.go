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
		app.Logger.Fatal("a payment id is required to create a refund")
	}

	var r mollie.Refund
	{
		r = mollie.Refund{
			Amount:      promptAmount(),
			Description: promptStringClean("refund description", ""),
			Metadata:    promptStringClean("additional metadata", ""),
		}
	}

	res, rs, err := app.API.Refunds.Create(context.Background(), payment, r, nil)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Refunds, rs, res)

	if verbose {
		app.Logger.Infof("refund for payment %s created", payment)
		app.Logger.Infof("request target: %s", rs.Links.Self.Href)
		app.Logger.Infof("request docs: %s", rs.Links.Documentation.Href)
	}

	disp := displayers.MollieRefund{Refund: rs}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(parseFieldsFromFlag(cmd, Refunds), refundsCols()),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
