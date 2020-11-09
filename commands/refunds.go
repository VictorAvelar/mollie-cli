package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	refundsCols = []string{
		"ID",
		"Payment",
		"Order",
		"Settlement",
		"Amount",
		"Status",
		"Description",
		"Created at",
	}
)

// Refunds builds the refunds commands tree.
func Refunds() *command.Command {
	r := command.Builder(
		nil,
		command.Config{
			Namespace: "refunds",
			Aliases:   []string{"refs", "rf"},
			ShortDesc: "All operations to handle refunds",
		},
		noCols,
	)

	lr := command.Builder(
		r,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieves refunds for the provided API token, or payment token",
			Example:   "mollie refunds list --payment=tr_test",
			Execute:   RunListRefunds,
		},
		noCols,
	)

	command.AddFlag(lr, command.FlagConfig{
		Name:  PaymentArg,
		Usage: "only Refunds for the specific Payment are returned",
	})

	command.AddFlag(lr, command.FlagConfig{
		Name:  FromArg,
		Usage: "offset the result set to the refund with this ID",
	})

	command.AddFlag(lr, command.FlagConfig{
		Name:  LimitArg,
		Usage: "the number of refunds to return (with a maximum of 250)",
	})

	command.AddFlag(lr, command.FlagConfig{
		Name:  EmbedArg,
		Usage: "embedding additional information (payments)",
	})

	gr := command.Builder(
		r,
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single Refund by its ID",
			LongDesc:  "Retrieve a single Refund by its ID. Note the Payment’s ID is needed as well",
			Example:   "mollie refunds get --id=rf_test --payment=tr_test",
			Execute:   RunGetRefund,
		},
		noCols,
	)

	command.AddFlag(gr, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the refund id/tokenx",
		Required: true,
	})

	command.AddFlag(gr, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "original payment id/token",
		Required: true,
	})

	command.AddFlag(gr, command.FlagConfig{
		Name:  EmbedArg,
		Usage: "embedding additional information (payments)",
	})

	return r
}

// RunListRefunds retrieves a list of refunds.
func RunListRefunds(cmd *cobra.Command, args []string) {
	var opts mollie.ListRefundOptions
	{
		opts.Embed = mollie.EmbedValue(ParseStringFromFlags(cmd, EmbedArg))
		opts.Limit = ParseStringFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
	}

	payment := ParseStringFromFlags(cmd, PaymentArg)

	if verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(LimitArg, opts.Limit)
		PrintNonemptyFlagValue(FromArg, opts.From)
		PrintNonemptyFlagValue(EmbedArg, string(opts.Embed))
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

	disp := displayers.MollieRefundList{RefundList: refunds}

	err = command.Display(refundsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetRefund retrieves a refund by its id/token and the payment id/token.
func RunGetRefund(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)
	payment := ParseStringFromFlags(cmd, PaymentArg)
	embed := mollie.EmbedValue(ParseStringFromFlags(cmd, EmbedArg))

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(EmbedArg, string(embed))
	}

	r, err := API.Refunds.Get(payment, id, &mollie.RefundOptions{Embed: embed})
	if err != nil {
		logger.Fatal(err)
	}

	disp := displayers.MollieRefund{
		Refund: &r,
	}

	err = command.Display(refundsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

func getRefundList(opts *mollie.ListRefundOptions, payment string) (*mollie.RefundList, error) {
	if payment != "" {
		return API.Refunds.ListRefundPayment(payment, opts)
	}

	return API.Refunds.ListRefund(opts)
}
