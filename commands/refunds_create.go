package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
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

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	rs, err := API.Refunds.Create(payment, r, nil)
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

	disp := displayers.MollieRefund{Refund: &rs}

	err = printer.Display(
		&disp,
		display.FilterColumns(parseFieldsFromFlag(cmd), refundsCols()),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
