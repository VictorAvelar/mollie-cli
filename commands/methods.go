package commands

import (
	"fmt"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/avocatl/admiral/pkg/prompter"
	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
)

var (
	methodsCols = []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"ISSUERS",
		"MIN_AMOUNT",
		"MAX_AMOUNT",
		"LOGO",
	}
)

// Methods builds the methods commands tree.
func Methods() *commander.Command {
	m := commander.Builder(nil, commander.Config{
		Namespace: "methods",
		Aliases:   []string{"vendors", "meths"},
		ShortDesc: "All payment methods that Mollie offers and can be activated",
	}, methodsCols)

	AddIncludeFlag(m, true)
	AddPrompterFlag(m, true)

	lm := commander.Builder(
		m,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieves all enabled payment methods",
			LongDesc: `Retrieves all enabled payment methods.

To check the payment method embeded resources use the get payment methods command.`,
			Example: "mollie methods list --locale=de_DE --sequence-type=recurring",
			Execute: RunListPaymentMethods,
		},
		methodsCols,
	)
	AddResourceFlag(lm)
	AddSequenceTypeFlag(lm)
	AddCurrencyFlags(lm)
	AddLocaleFlag(lm)
	AddBillingCountryFlag(lm)
	AddWalletFlag(lm)

	ga := commander.Builder(
		m,
		commander.Config{
			Namespace: "all",
			ShortDesc: "Retrieve all payment methods that Mollie offers and can be activated by the Organization.",
			LongDesc: `Retrieve all payment methods that Mollie offers and can be activated by the Organization.
The results are not paginated. New payment methods can be activated via the Enable payment method
endpoint in the Profiles API.

To check the payment method embeded resources use the get payment methods command.`,
			Execute: RunGetAllMethods,
			Example: "mollie methods all --locale=nl_NL",
		},
		methodsCols,
	)

	AddLocaleFlag(ga)
	AddCurrencyFlags(ga)

	gm := commander.Builder(
		m,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single method by its ID.",
			LongDesc: `Retrieve a single method by its ID. Note that if a method is not available on the website profile
a status 404 Not found is returned. When the method is not enabled,a status 403 Forbidden
is returned. You can enable payments methods via the Enable payment method endpoint in the
Profiles API, or via your Mollie Dashboard.`,
			Execute: RunGetPaymentMethods,
			Example: "mollie methods get --id=creditcard --locale=pt_PT",
		},
		methodsCols,
	)

	AddIDFlag(gm, true)
	AddLocaleFlag(gm)
	AddCurrencyCodeFlag(gm)

	return m
}

// RunListPaymentMethods retrieves all the payment methods enabled for the token.
func RunListPaymentMethods(cmd *cobra.Command, args []string) {
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

// RunGetAllMethods retrieves all available payment methods for the token.
func RunGetAllMethods(cmd *cobra.Command, args []string) {
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
		PrintJsonP(m)
	}

	if verbose {
		logger.Infof("received %d payment methods", m.Count)
		logger.Infof("request performed: %s", m.Links.Self.Href)
		logger.Infof("documentation: %s", m.Links.Documentation.Href)
	}

	disp := &displayers.MollieListMethods{ListMethods: m}

	err = printer.Display(disp)
	if err != nil {
		logger.Fatal(err)
	}
}

type getMethodPropmter struct {
	Locale  string
	Include string
}

// RunGetPaymentMethods retrieves a payment method by its id.
func RunGetPaymentMethods(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetString(IDArg)
	if err != nil {
		logger.Fatal(err)
	}

	var opts mollie.MethodsOptions
	{
		if ParsePromptBool(cmd) {
			v, err := prompter.Struct(&getMethodPropmter{})
			if err != nil {
				logger.Fatal(err)
			}
			val := v.(*getMethodPropmter)

			opts.Locale = mollie.Locale(val.Locale)
			opts.Include = val.Include
		} else {
			opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
			opts.Currency = ParseStringFromFlags(cmd, CurrencyArg)
			opts.Include = ParseStringFromFlags(cmd, IncludeArg)

		}
	}

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	m, err := API.Methods.Get(id, &opts)
	if err != nil {
		logger.Fatal(err)
	}

	if json {
		PrintJsonP(m)
	}

	err = printer.DisplayMany(getMethodsDisplayables(m))
	if err != nil {
		logger.Fatal(err)
	}
}

func getMethodsDisplayables(m *mollie.PaymentMethodInfo) []display.Displayable {
	method := &displayers.MollieMethod{
		PaymentMethodInfo: m,
	}

	var dp []display.Displayable
	{
		dp = append(dp, method)

		if verbose {
			dp = append(dp, displayers.NewSimpleTextDisplayer("=", fmt.Sprintf("Embeded issuers for method %s", m.Description)))
		}

		if len(m.Issuers) > 0 {
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

		colmap := displayers.NewSimpleTextDisplayer("=", text)

		dp = append([]display.Displayable{colmap}, dp...)
	}

	return dp
}
