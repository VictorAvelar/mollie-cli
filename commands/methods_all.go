package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
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

To check the payment method embedded resources use the get payment methods command.`,
			Execute: getAllMethodsAction,
			Example: "mollie methods all --locale=nl_NL",
		},
		getMethodsCols(),
	)

	AddLocaleFlag(ga)
	AddCurrencyFlags(ga)
}

func getAllMethodsAction(cmd *cobra.Command, args []string) {
	var opts mollie.PaymentMethodsListOptions
	{
		opts.Currency = ParseStringFromFlags(cmd, AmountCurrencyArg)
		opts.AmountValue = ParseStringFromFlags(cmd, AmountValueArg)
		opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
		opts.Include = ParseStringFromFlags(cmd, IncludeArg)
	}

	res, m, err := app.API.PaymentMethods.All(context.Background(), &opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Methods, m, res)

	if verbose {
		app.Logger.Infof("received %d payment methods", m.Count)
		app.Logger.Infof("request performed: %s", m.Links.Self.Href)
		app.Logger.Infof("documentation: %s", m.Links.Documentation.Href)
	}

	disp := &displayers.MollieListMethods{PaymentMethodsList: m}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Methods),
			getMethodsCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
