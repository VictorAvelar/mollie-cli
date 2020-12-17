package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
)

var (
	methodsCols = []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"MIN_AMOUNT",
		"MAX_AMOUNT",
		"LOGO",
	}
)

// Methods builds the methods commands tree.
func Methods() *command.Command {
	m := command.Builder(
		nil,
		command.Config{
			Namespace: "methods",
			Aliases:   []string{"meth", "vendors"},
			ShortDesc: "All payment methods that Mollie offers and can be activated",
		},
		methodsCols,
	)

	lm := command.Builder(
		m,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieves all enabled payment methods",
			Execute:   RunListPaymentMethods,
			Example:   "mollie methods list --locale=de_DE --sequence-type=recurring",
		},
		methodsCols,
	)
	command.AddFlag(lm, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "get the payment method name in the corresponding language",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  SequenceTypeArg,
		Usage: "filter methods by sequence type (oneoff, first, recurring)",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  AmountCurrencyArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-value)",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  AmountValueArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-currency)",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  ResourceArg,
		Usage: "filter for methods that can be used in combination with the provided resource (orders/payments)",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  BillingCountryArg,
		Usage: "filter for methods supporting the ISO-3166 alpha-2 customer billing country",
	})
	command.AddFlag(lm, command.FlagConfig{
		Name:  WalletsArg,
		Usage: "a comma-separated list of the wallets you support in your checkout (applepay)",
	})

	ga := command.Builder(
		m,
		command.Config{
			Namespace: "all",
			ShortDesc: "Retrieve all payment methods that Mollie offers and can be activated by the Organization.",
			LongDesc: `Retrieve all payment methods that Mollie offers and can be activated by the Organization. 
The results are not paginated. New payment methods can be activated via the Enable payment method 
endpoint in the Profiles API.`,
			Execute: RunGetAllMethods,
			Example: "mollie methods all --locale=nl_NL",
		},
		methodsCols,
	)

	command.AddFlag(ga, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "get the payment method name in the corresponding language",
	})
	command.AddFlag(ga, command.FlagConfig{
		Name:  AmountCurrencyArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-value)",
	})
	command.AddFlag(ga, command.FlagConfig{
		Name:  AmountValueArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-currency)",
	})

	gm := command.Builder(
		m,
		command.Config{
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

	command.AddFlag(gm, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment method id",
		Required: true,
	})
	command.AddFlag(gm, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "get the payment method name in the corresponding language",
	})
	command.AddFlag(gm, command.FlagConfig{
		Name:  CurrencyArg,
		Usage: "the currency to receiving the minimumAmount and maximumAmount in",
	})

	return m
}

// RunListPaymentMethods retrieves all the payment methods enabled for the token.
func RunListPaymentMethods(cmd *cobra.Command, args []string) {
	var opts mollie.MethodsOptions
	{
		opts.SequenceType = mollie.SequenceType(ParseStringFromFlags(cmd, SequenceTypeArg))
		opts.AmountCurrency = ParseStringFromFlags(cmd, AmountCurrencyArg)
		opts.AmountValue = ParseStringFromFlags(cmd, AmountValueArg)
		opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
		opts.IncludeWallets = ParseStringFromFlags(cmd, WalletsArg)
		opts.Resource = ParseStringFromFlags(cmd, ResourceArg)
		opts.BillingCountry = ParseStringFromFlags(cmd, BillingCountryArg)
	}

	if verbose {
		PrintNonemptyFlagValue(SequenceTypeArg, string(opts.SequenceType))
		PrintNonemptyFlagValue(LocaleArg, string(opts.Locale))
		PrintNonemptyFlagValue(WalletsArg, opts.IncludeWallets)
		PrintNonemptyFlagValue(ResourceArg, opts.Resource)
		PrintNonemptyFlagValue(AmountCurrencyArg, opts.AmountCurrency)
		PrintNonemptyFlagValue(AmountValueArg, opts.AmountValue)
		PrintNonemptyFlagValue(BillingCountryArg, opts.BillingCountry)
	}

	ms, err := API.Methods.List(&opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("received %d payment methods", ms.Count)
		logger.Infof("request performed: %s", ms.Links.Self.Href)
		logger.Infof("documentation: %s", ms.Links.Documentation.Href)
	}

	disp := displayers.MollieListMethods{
		ListMethods: ms,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), methodsCols),
		disp.KV(),
	)
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
	}

	if verbose {
		PrintNonemptyFlagValue(LocaleArg, string(opts.Locale))
		PrintNonemptyFlagValue(AmountCurrencyArg, opts.AmountCurrency)
		PrintNonemptyFlagValue(AmountValueArg, opts.AmountValue)
	}

	m, err := API.Methods.All(&opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("received %d payment methods", m.Count)
		logger.Infof("request performed: %s", m.Links.Self.Href)
		logger.Infof("documentation: %s", m.Links.Documentation.Href)
	}

	disp := &displayers.MollieListMethods{ListMethods: m}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), methodsCols),
		disp.KV(),
	)
	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetPaymentMethods retrieves a payment method by its id.
func RunGetPaymentMethods(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetString(IDArg)
	if err != nil {
		logger.Fatal(err)
	}

	var opts mollie.MethodsOptions
	{
		opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
		opts.Currency = ParseStringFromFlags(cmd, CurrencyArg)
	}

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
		PrintNonemptyFlagValue(LocaleArg, string(opts.Locale))
		PrintNonemptyFlagValue(CurrencyArg, opts.Currency)
	}

	m, err := API.Methods.Get(id, &opts)
	if err != nil {
		logger.Fatal(err)
	}

	disp := &displayers.MollieMethod{
		PaymentMethodInfo: m,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), methodsCols),
		disp.KV(),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
