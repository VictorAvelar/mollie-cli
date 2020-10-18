package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
)

var (
	methodsCols = []string{
		"ID",
		"Name",
		"Minimum Amount",
		"Maximum Amount",
	}
)

// Methods builds the methods commands tree.
func Methods() *command.Command {
	m := &command.Command{
		Command: &cobra.Command{
			Use:   "methods",
			Short: "All payment methods that Mollie offers and can be activated",
		},
	}

	lm := command.Builder(
		m,
		"list",
		"Retrieve all enabled payment methods",
		`For test mode, payment methods are returned that are enabled in the Dashboard 
(or the activation is pending).

For live mode, payment methods are returned that have been activated on your 
account and have been enabled in the Dashboard.
`,
	RunListPaymentMethods,
	methodsCols,
	)

	command.AddStringFlag(lm, LocaleArg, "", "", "get the payment method name in the corresponding language", false)
	command.AddStringFlag(lm, SequenceTypeArg, "", "", "filter methods by sequence type (oneoff, first, recurring)", false)
	command.AddStringFlag(lm, AmountCurrencyArg, "", "", "get only payment methods that support the amount and currency (linked to amount-value)", false)
	command.AddStringFlag(lm, AmountValueArg, "", "", "get only payment methods that support the amount and currency (linked to amount-currency)", false)
	command.AddStringFlag(lm, ResourceArg, "", "", "filter for methods that can be used in combination with the provided resource (orders/payments)", false)
	command.AddStringFlag(lm, BillingCountryArg, "", "", "filter for methods supporting the ISO-3166 alpha-2 customer billing country", false)
	command.AddStringFlag(lm, WalletsArg, "", "", "a comma-separated list of the wallets you support in your checkout (applepay)", false)

	ga := command.Builder(
		m,
		"all",
		"Retrieve all payment methods that Mollie offers and can be activated by the Organization.",
		`Retrieve all payment methods that Mollie offers and can be activated by the Organization. 
The results are not paginated. New payment methods can be activated via the Enable payment method 
endpoint in the Profiles API.`,
		RunGetAllMethods,
		methodsCols,
	)

	command.AddStringFlag(ga, LocaleArg, "", "", "get the payment method name in the corresponding language", false)
	command.AddStringFlag(ga, AmountCurrencyArg, "", "", "get only payment methods that support the amount and currency (linked to amount-value)", false)
	command.AddStringFlag(ga, AmountValueArg, "", "", "get only payment methods that support the amount and currency (linked to amount-currency)", false)

	gm := command.Builder(
		m,
		"get",
		"Retrieve a single method by its ID.",
		`Retrieve a single method by its ID. Note that if a method is not available on the website profile 
a status 404 Not found is returned. When the method is not enabled,a status 403 Forbidden 
is returned. You can enable payments methods via the Enable payment method endpoint in the 
Profiles API, or via your Mollie Dashboard.`,
		RunGetPaymentMethods,
		methodsCols,
	)

	command.AddStringFlag(gm, IDArg, "", "", "the payment method id", true)
	command.AddStringFlag(gm, LocaleArg, "", "", "get the payment method name in the corresponding language", false)
	command.AddStringFlag(gm, CurrencyArg, "", "", "the currency to receiving the minimumAmount and maximumAmount in", false)

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

	if Verbose {
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
		logrus.Fatal(err)
	}

	if Verbose {
		logrus.Infof("received %d payment methods", ms.Count)
		logrus.Infof("request performed: %s", ms.Links.Self.Href)
		logrus.Infof("documentation: %s", ms.Links.Docs.Href)
	}

	lpm := displayers.MollieListMethods{
		ListMethods: ms,
	}

	err = command.Display(methodsCols, lpm.KV())
	if err != nil {
		logrus.Fatal(err)
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

	if Verbose {
		PrintNonemptyFlagValue(LocaleArg, string(opts.Locale))
		PrintNonemptyFlagValue(AmountCurrencyArg, opts.AmountCurrency)
		PrintNonemptyFlagValue(AmountValueArg, opts.AmountValue)
	}

	m, err := API.Methods.All(&opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if Verbose {
		logrus.Infof("received %d payment methods", m.Count)
		logrus.Infof("request performed: %s", m.Links.Self.Href)
		logrus.Infof("documentation: %s", m.Links.Docs.Href)
	}

	mdis := &displayers.MollieListMethods{ListMethods: m}

	err = command.Display(methodsCols, mdis.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}

// RunGetPaymentMethods retrieves a payment method by its id.
func RunGetPaymentMethods(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetString(IDArg)
	if err != nil {
		log.Fatal(err)
	}

	var opts mollie.MethodsOptions
	{
		opts.Locale = mollie.Locale(ParseStringFromFlags(cmd, LocaleArg))
		opts.Currency = ParseStringFromFlags(cmd, CurrencyArg)
	}

	if Verbose {
		PrintNonemptyFlagValue(IDArg, id)
		PrintNonemptyFlagValue(LocaleArg, string(opts.Locale))
		PrintNonemptyFlagValue(CurrencyArg, opts.Currency)
	}

	m, err := API.Methods.Get(id, &opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mdis := &displayers.MollieMethod{
		PaymentMethodInfo: m,
	}

	err = command.Display(methodsCols, mdis.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}
