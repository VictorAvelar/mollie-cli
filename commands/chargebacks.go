package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

// Chargebacks creates the chargebacks commands tree.
func Chargebacks() *command.Command {
	cb := command.Builder(
		nil,
		command.Config{
			Namespace: "chargebacks",
			ShortDesc: "Operations with the Chargebacks API",
			Aliases:   []string{"cb", "cback"},
		},
		noCols,
	)

	gcb := command.Builder(
		cb,
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.",
			LongDesc: `Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.
A debit to a depositor's account for an item that has been previously credited, as for a returned bad check.`,
			Execute: RunGetChargebacks,
		},
		noCols,
	)
	command.AddFlag(gcb, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "original payment id/token",
		Required: true,
	})
	command.AddFlag(gcb, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the chargeback id/token",
		Required: true,
	})
	command.AddFlag(gcb, command.FlagConfig{
		Name:  EmbedArg,
		Usage: "a comma separated list of embeded resources",
	})

	lcb := command.Builder(
		cb,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all received chargebacks",
			LongDesc: `Retrieve all received chargebacks. If the payment-specific endpoint is used, only chargebacks 
for that specific payment are returned.`,
			Execute: RunListChargebacks,
		},
		noCols,
	)
	command.AddFlag(lcb, command.FlagConfig{
		Name:  EmbedArg,
		Usage: "a comma separated list of embeded resources",
	})

	return cb
}

// RunGetChargebacks will retrieve all the received chargebacks
// for a payment.
func RunGetChargebacks(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	chargeback := ParseStringFromFlags(cmd, IDArg)
	embed := ParseStringFromFlags(cmd, EmbedArg)

	if verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(IDArg, chargeback)
		PrintNonemptyFlagValue(EmbedArg, embed)
	}

	cb, err := API.Chargebacks.Get(payment, chargeback, nil)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
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

	if verbose {
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

	if verbose {
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
