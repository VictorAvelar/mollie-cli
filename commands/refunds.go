package commands

import (
	"os"

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

	cr := command.Builder(
		r,
		command.Config{
			Namespace: "create",
			Aliases:   []string{"new", "add"},
			ShortDesc: "Creates a Refund on the Payment",
			LongDesc:  "Creates a Refund on the Payment. The refunded amount is credited to your customer.",
			Example:   "mollie refunds create --payment=tr_test",
			Execute:   RunCreateRefund,
		},
		noCols,
	)

	command.AddFlag(cr, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "original payment id/token",
		Required: true,
	})

	command.AddFlag(cr, command.FlagConfig{
		Name:     AmountValueArg,
		Usage:    "a string containing the exact amount you want to refund",
		Required: true,
	})

	command.AddFlag(cr, command.FlagConfig{
		Name:     AmountCurrencyArg,
		Usage:    "an ISO 4217 currency code (same as payment)",
		Required: true,
	})

	command.AddFlag(cr, command.FlagConfig{
		Name:  DescriptionArg,
		Usage: "the description of the refund you are creating",
	})

	command.AddFlag(cr, command.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like to attach to the refund",
	})

	dr := command.Builder(
		r,
		command.Config{
			Namespace: "cancel",
			Aliases:   []string{"delete", "remove", "cncl"},
			ShortDesc: "for certain payment methods where cancelation is possible.",
			LongDesc: `For certain payment methods, like iDEAL, the underlying banking system will delay refunds
until the next day. Until that time, refunds may be canceled manually in the Mollie Dashboard, 
or programmatically by using this endpoint.

A Refund can only be canceled while its status field is either queued or pending.`,
			Example: "mollie refunds cancel --id=rf_test --payment=tr_test",
			Execute: RunCancelRefund,
		},
		noCols,
	)

	command.AddFlag(dr, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the refund id/token",
		Required: true,
	})

	command.AddFlag(dr, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "original payment id/token",
		Required: true,
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

// RunCreateRefund creates a refund for the given payment.
func RunCreateRefund(cmd *cobra.Command, args []string) {
	r := mollie.Refund{}
	r.Amount = &mollie.Amount{
		Currency: ParseStringFromFlags(cmd, AmountCurrencyArg),
		Value:    ParseStringFromFlags(cmd, AmountValueArg),
	}
	r.Description = ParseStringFromFlags(cmd, DescriptionArg)
	r.Metadata = ParseStringFromFlags(cmd, MetadataArg)

	payment := ParseStringFromFlags(cmd, PaymentArg)

	if verbose {
		PrintNonemptyFlagValue(AmountCurrencyArg, r.Amount.Currency)
		PrintNonemptyFlagValue(AmountValueArg, r.Amount.Value)
		PrintNonemptyFlagValue(DescriptionArg, r.Description)
		PrintNonemptyFlagValue(MetadataArg, r.Metadata.(string))
		PrintNonemptyFlagValue(PaymentArg, payment)
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

	err = command.Display(refundsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunCancelRefund cancels a refund for allowed payment methods
// and when the status is queued or pending.
func RunCancelRefund(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
		PrintNonemptyFlagValue(PaymentArg, payment)
	}

	err := API.Refunds.Cancel(payment, id, nil)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("refund %s cancelled", id)

	os.Exit(0)
}

func getRefundList(opts *mollie.ListRefundOptions, payment string) (*mollie.RefundList, error) {
	if payment != "" {
		return API.Refunds.ListRefundPayment(payment, opts)
	}

	return API.Refunds.ListRefund(opts)
}
