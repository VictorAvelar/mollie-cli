package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/avocatl/admiral/pkg/prompter"
	"github.com/spf13/cobra"
)

func listPaymentMethodsCmd(p *commander.Command) {
	lm := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieves all enabled payment methods",
			LongDesc: `Retrieves all enabled payment methods.

To check the payment method embedded resources use the get payment methods command.`,
			Example: "mollie methods list --locale=de_DE --sequence-type=recurring",
			Execute: listPaymentMethodsAction,
		},
		getMethodsCols(),
	)
	AddResourceFlag(lm)
	AddSequenceTypeFlag(lm)
	AddCurrencyFlags(lm)
	AddLocaleFlag(lm)
	AddBillingCountryFlag(lm)
	AddWalletFlag(lm)
}

func listPaymentMethodsAction(cmd *cobra.Command, args []string) {
	var opts mollie.PaymentMethodsListOptions
	{
		if ParsePromptBool(cmd) {
			oi, err := prompter.Struct(&opts)
			if err != nil {
				logger.Fatal(err)
			}

			optsi := oi.(*mollie.PaymentMethodsListOptions)
			opts = *optsi
		} else {
			opts.SequenceType = mollie.SequenceType(ParseStringFromFlags(cmd, SequenceTypeArg))
			opts.AmountCurrency = ParseStringFromFlags(cmd, AmountCurrencyArg)
			opts.AmountValue = ParseStringFromFlags(cmd, AmountValueArg)
			opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
			opts.IncludeWallets = ParseStringFromFlags(cmd, WalletsArg)
			opts.Resource = ParseStringFromFlags(cmd, ResourceArg)
			opts.BillingCountry = ParseStringFromFlags(cmd, BillingCountryArg)
			opts.Include = ParseStringFromFlags(cmd, IncludeArg)
		}

	}

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	_, ms, err := API.PaymentMethods.List(context.Background(), &opts)
	if err != nil {
		logger.Fatal(err)
	}

	if json {
		printJSONP(ms)
	}

	if verbose {
		logger.Infof("received %d payment methods", ms.Count)
		logger.Infof("request performed: %s", ms.Links.Self.Href)
		logger.Infof("documentation: %s", ms.Links.Documentation.Href)
	}

	disp := displayers.MollieListMethods{
		PaymentMethodsList: ms,
	}

	err = printer.Display(&disp, display.FilterColumns(
		parseFieldsFromFlag(cmd), getMethodsCols(),
	))

	if err != nil {
		logger.Fatal(err)
	}
}
