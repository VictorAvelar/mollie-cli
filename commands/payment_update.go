package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func updatePaymentCmd(p *commander.Command) *commander.Command {
	up := commander.Builder(
		p,
		commander.Config{
			Namespace: "update",
			Aliases:   []string{"edit", "change", "mutate"},
			ShortDesc: "Update some details of a created payment",
			LongDesc: `Updates basic payment information, for a more advanced
and complete update workflow check the prompt subcommand.`,
			Execute: updatePaymentAction,
			Example: "mollie payments update --id=test_token --description=updated",
		},
		getPaymentCols(),
	)

	addUpdatePaymentFlags(up)
	updatePaymentPromptCmd(up)

	return up
}

func addUpdatePaymentFlags(up *commander.Command) {
	commander.AddFlag(up, commander.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment token/id",
		Required: true,
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  DescriptionArg,
		Usage: "the description of the payment to show to your customers when possible",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  RedirectURLArg,
		Usage: "the URL your customer will be redirected to after the payment process",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  WebhookURLArg,
		Usage: "set the webhook URL, where we will send payment status updates to",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  MetadataArg,
		Usage: "any data you like, for example a string or a JSON object",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  MethodArg,
		Usage: "change the payment to a different payment method.",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  LocaleArg,
		Usage: "update the language to be used in the hosted payment pages",
	})
	commander.AddFlag(up, commander.FlagConfig{
		Name:  RPMToCountryArg,
		Usage: "parameter to restrict the payment methods available to your customer to those from a single country",
	})
}

func updatePaymentAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)
	desc := ParseStringFromFlags(cmd, DescriptionArg)
	rURL := ParseStringFromFlags(cmd, RedirectURLArg)
	whURL := ParseStringFromFlags(cmd, WebhookURLArg)
	meta := ParseStringFromFlags(cmd, MetadataArg)
	method := ParseStringFromFlags(cmd, MethodArg)
	locale := ParseStringFromFlags(cmd, LocaleArg)
	rpmCountry := ParseStringFromFlags(cmd, RPMToCountryArg)

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	l := mollie.Locale(locale)
	c := mollie.Locale(rpmCountry)
	m := mollie.PaymentMethod(method)

	_, p, err := API.Payments.Update(context.Background(), id, mollie.Payment{
		RedirectURL:                     rURL,
		Description:                     desc,
		WebhookURL:                      whURL,
		Metadata:                        meta,
		Locale:                          l,
		RestrictPaymentMethodsToCountry: c,
		Method:                          m,
	})
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
		logger.Infof("payment successfully updated")
	}

	disp := displayers.MolliePayment{Payment: p}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			getPaymentCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
