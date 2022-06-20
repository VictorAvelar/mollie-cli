package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func refundsGetCmd(p *commander.Command) *commander.Command {
	gr := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single Refund by its ID",
			LongDesc:  "Retrieve a single Refund by its ID. Note the Payment’s ID is needed as well",
			Example:   "mollie refunds get --id=rf_test --payment=tr_test",
			Execute:   getRefundAction,
		},
		refundsCols(),
	)

	AddIDFlag(gr, true)
	AddPaymentFlag(gr)
	AddEmbedFlag(gr)

	return gr
}

func getRefundAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)
	payment := ParseStringFromFlags(cmd, PaymentArg)
	embed := mollie.EmbedValue(ParseStringFromFlags(cmd, EmbedArg))

	var opts *mollie.RefundOptions
	{
		if embed != "" {
			opts = &mollie.RefundOptions{Embed: embed}
		}
	}

	res, r, err := app.API.Refunds.Get(context.Background(), payment, id, opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Refunds, r, res)

	disp := displayers.MollieRefund{
		Refund: r,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(parseFieldsFromFlag(cmd, Refunds), refundsCols()),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
