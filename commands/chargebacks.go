package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/VictorAvelar/mollie-cli/internal/runners"
	"github.com/spf13/cobra"
)

// Chargebacks creates the chargebacks commands tree.
func Chargebacks() *command.Command {
	cb := command.Builder(
		nil,
		"chargebacks",
		"Operations with the Chargebacks API",
		``,
		runners.NopRunner,
		[]string{},
	)
	cb.Aliases = []string{"cb", "cback"}

	gcb := command.Builder(
		cb,
		"get",
		"Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.",
		`Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.
A debit to a depositor's account for an item that has been previously credited, as for a returned bad check.`,
		RunGetChargebacks,
		[]string{},
	)

	command.AddStringFlag(gcb, PaymentArg, "", "", "original payment id/token", true)
	command.AddStringFlag(gcb, IDArg, "", "", "the chargeback id", true)
	command.AddStringFlag(gcb, EmbedArg, "", "", "a comma separated list of embeded resources", false)

	lcb := command.Builder(
		cb,
		"list",
		"Retrieve all received chargebacks",
		`Retrieve all received chargebacks. If the payment-specific endpoint is used, only chargebacks 
for that specific payment are returned.`,
		RunListChargebacks,
		[]string{},
	)

	command.AddStringFlag(lcb, EmbedArg, "", "", "a comma separated list of embeded resources", false)

	return cb
}

// RunGetChargebacks will retrieve all the received chargebacks
// for a payment.
func RunGetChargebacks(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	chargeback := ParseStringFromFlags(cmd, IDArg)
	embed := ParseStringFromFlags(cmd, EmbedArg)

	if Verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(IDArg, chargeback)
		PrintNonemptyFlagValue(EmbedArg, embed)
	}

	cb, err := API.Chargebacks.Get(payment, chargeback, nil)
	if err != nil {
		logger.Fatal(err)
	}

	if Verbose {
		logger.Infof("request target: %s", cb.Links.Self.Href)
		logger.Infof("request docs: %s", cb.Links.Documentation.Href)
	}

	display := &displayers.MollieChargeback{Chargeback: &cb}

	err = command.Display([]string{"ID", "Payment", "Amount", "Settlement", "Created at"}, display.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunListChargebacks will retrieve all the chargebacks for the
// current token.
func RunListChargebacks(cmd *cobra.Command, args []string) {
	embed := ParseStringFromFlags(cmd, EmbedArg)

	if Verbose {
		PrintNonemptyFlagValue(EmbedArg, embed)
	}

	var opt mollie.ListChargebackOptions
	{
		opt.Embed = embed
	}

	cbs, err := API.Chargebacks.List(&opt)
	if err != nil {
		logger.Fatal(err)
	}

	if Verbose {
		logger.Infof("response with %d chargebacks", cbs.Count)
		logger.Infof("request target: %s", cbs.Links.Self.Href)
		logger.Infof("request docs: %s", cbs.Links.Docs.Href)
	}

	display := displayers.MollieChargebackList{ChargebackList: cbs}

	err = command.Display([]string{"ID", "Payment", "Amount", "Settlement", "Created at"}, display.KV())
	if err != nil {
		logger.Fatal(err)
	}
}
