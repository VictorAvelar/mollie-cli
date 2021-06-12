package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func allPaymentMethodsCmd(p *commander.Command) {
	ga := commander.Builder(
		p,
		commander.Config{
			Namespace: "all",
			ShortDesc: "Retrieve all payment methods that Mollie offers and can be activated by the Organization.",
			LongDesc: `Retrieve all payment methods that Mollie offers and can be activated by the Organization.
The results are not paginated. New payment methods can be activated via the Enable payment method
endpoint in the Profiles API.

To check the payment method embeded resources use the get payment methods command.`,
			Execute: getAllMethodsAction,
			Example: "mollie methods all --locale=nl_NL",
		},
		getMethodsCols(),
	)

	AddLocaleFlag(ga)
	AddCurrencyFlags(ga)
}

func getAllMethodsAction(cmd *cobra.Command, args []string) {
	var opts mollie.MethodsOptions
	{
		opts.AmountCurrency = ParseStringFromFlags(cmd, AmountCurrencyArg)
		opts.AmountValue = ParseStringFromFlags(cmd, AmountValueArg)
		opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
		opts.Include = ParseStringFromFlags(cmd, IncludeArg)
	}

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	m, err := API.Methods.All(&opts)
	if err != nil {
		logger.Fatal(err)
	}

	if json {
		printJSONP(m)
	}

	if verbose {
		logger.Infof("received %d payment methods", m.Count)
		logger.Infof("request performed: %s", m.Links.Self.Href)
		logger.Infof("documentation: %s", m.Links.Documentation.Href)
	}

	disp := &displayers.MollieListMethods{ListMethods: m}

	err = printer.Display(disp, display.FilterColumns(
		parseFieldsFromFlag(cmd), getMethodsCols(),
	))
	if err != nil {
		logger.Fatal(err)
	}
}
