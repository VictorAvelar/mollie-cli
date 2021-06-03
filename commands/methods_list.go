package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
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

To check the payment method embeded resources use the get payment methods command.`,
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
	var opts mollie.MethodsOptions
	{
		if ParsePromptBool(cmd) {
			oi, err := prompter.Struct(&opts)
			if err != nil {
				logger.Fatal(err)
			}

			optsi := oi.(*mollie.MethodsOptions)
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

	ms, err := API.Methods.List(&opts)
	if err != nil {
		logger.Fatal(err)
	}

	if json {
		PrintJsonP(ms)
	}

	if verbose {
		logger.Infof("received %d payment methods", ms.Count)
		logger.Infof("request performed: %s", ms.Links.Self.Href)
		logger.Infof("documentation: %s", ms.Links.Documentation.Href)
	}

	disp := displayers.MollieListMethods{
		ListMethods: ms,
	}

	err = printer.Display(&disp)

	if err != nil {
		logger.Fatal(err)
	}
}
