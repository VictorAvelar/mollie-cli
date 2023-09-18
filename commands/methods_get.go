package commands

import (
	"context"
	"fmt"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/avocatl/admiral/pkg/prompter"
	"github.com/spf13/cobra"
)

func getPaymentMethodCmd(p *commander.Command) {
	gm := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single method by its ID.",
			LongDesc: `Retrieve a single method by its ID. Note that if a method is not available on the website profile
a status 404 Not found is returned. When the method is not enabled,a status 403 Forbidden
is returned. You can enable payments methods via the Enable payment method endpoint in the
Profiles API, or via your Mollie Dashboard.`,
			Execute:  getPaymentMethodAction,
			Example:  "mollie methods get --id=creditcard --locale=pt_PT",
			PostHook: printJSONAction,
		},
		getMethodsCols(),
	)

	AddIDFlag(gm, true)
	AddLocaleFlag(gm)
	AddCurrencyCodeFlag(gm)
}

type getMethodPrompter struct {
	Locale  string
	Include string
}

func getPaymentMethodAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	var opts mollie.PaymentMethodOptions
	{
		if ParsePromptBool(cmd) {
			v, err := prompter.Struct(&getMethodPrompter{})
			if err != nil {
				app.Logger.Fatal(err)
			}

			val, ok := v.(*getMethodPrompter)
			if !ok {
				display.Text("x", "error asserting method prompter")
			}

			opts.Locale = mollie.Locale(val.Locale)
			opts.Include = val.Include
		} else {
			opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
			opts.Currency = ParseStringFromFlags(cmd, CurrencyArg)
			opts.Include = ParseStringFromFlags(cmd, IncludeArg)
		}
	}

	res, m, err := app.API.PaymentMethods.Get(
		context.Background(),
		mollie.PaymentMethod(id),
		&opts,
	)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Methods, m, res)

	err = app.Printer.DisplayMany(
		getMethodsDisplayables(m),
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Methods),
			getMethodsCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func getMethodsDisplayables(m *mollie.PaymentMethodDetails) []display.Displayable {
	method := &displayers.MollieMethod{
		PaymentMethodDetails: m,
	}

	var dp []display.Displayable
	{
		dp = append(dp, method)

		if len(m.Issuers) > 0 {
			if verbose {
				dp = append(dp, display.Text(
					"=",
					fmt.Sprintf("Embedded issuers for method %s", m.Description),
				))
			}
			dp = append(dp, &displayers.MollieListPaymentMethodsIssuers{
				Issuers: m.Issuers,
			})
		}
	}

	if verbose {
		var text string
		{
			for c, d := range method.ColMap() {
				text += fmt.Sprintf("%s:\t%v\n", c, d)
			}
		}

		colmap := display.Text("=", text)

		dp = append([]display.Displayable{colmap}, dp...)
	}

	return dp
}
