package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func createRefundCmd(p *commander.Command) *commander.Command {
	cr := commander.Builder(
		p,
		commander.Config{
			Namespace: "create",
			Aliases:   []string{"new", "add"},
			ShortDesc: "Creates a Refund on the Payment",
			LongDesc:  "Creates a Refund on the Payment. The refunded amount is credited to your customer.",
			Example:   "mollie refunds create --payment=tr_test",
			Execute:   createRefundAction,
			PostHook:  printJsonAction,
		},
		refundsCols(),
	)

	AddPaymentFlag(cr)
	AddCurrencyFlags(cr)

	commander.AddFlag(cr, commander.FlagConfig{
		Name:  DescriptionArg,
		Usage: "the description of the refund you are creating",
	})

	commander.AddFlag(cr, commander.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like to attach to the refund",
	})

	promptCreateRefundCmd(cr)

	return cr
}

func createRefundAction(cmd *cobra.Command, args []string) {
	r := mollie.Refund{}
	r.Amount = &mollie.Amount{
		Currency: ParseStringFromFlags(cmd, AmountCurrencyArg),
		Value:    ParseStringFromFlags(cmd, AmountValueArg),
	}
	r.Description = ParseStringFromFlags(cmd, DescriptionArg)
	r.Metadata = ParseStringFromFlags(cmd, MetadataArg)

	payment := ParseStringFromFlags(cmd, PaymentArg)

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
