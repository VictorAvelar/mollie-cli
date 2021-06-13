package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func createPaymentCmd(p *commander.Command) *commander.Command {
	cpp := commander.Builder(
		p,
		commander.Config{
			Namespace: "create",
			ShortDesc: "Create a new payment",
			LongDesc: `Creates a new payment.
Description, value, currency and redirect-url are required values.`,
			Execute: createPaymentAction,
			Aliases: []string{"new", "start"},
			Example: "mollie payments create --amount-value=200.00 --amount-currency=USD --redirect-to=https://victoravelar.com --description='custom example payment'",
		},
		getPaymentCols(),
	)

	addCreatePaymentFlags(cpp)
	promptPaymentCmd(cpp)

	return cpp
}

func addCreatePaymentFlags(cpp *commander.Command) {
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:     DescriptionArg,
		Usage:    "the description of the payment you’re creating to be show to your customers when possible",
		Required: true,
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:     RedirectURLArg,
		Usage:    "the URL your customer will be redirected to after the payment process",
		Required: true,
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:     AmountCurrencyArg,
		Usage:    "get only payment methods that support the amount and currency (linked to amount-value)",
		Required: true,
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:     AmountValueArg,
		Usage:    "get only payment methods that support the amount and currency (linked to amount-currency)",
		Required: true,
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		FlagType: commander.BoolFlag,
		Name:     CancelableArg,
		Usage:    "indicates if the payment can be cancelled",
		Default:  true,
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  WebhookURLArg,
		Usage: "set the webhook URL, where we will send payment status updates to",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  MetadataArg,
		Usage: "any data you like, for example a string or a JSON object",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  MethodArg,
		Usage: "change the payment to a different payment method.",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  LocaleArg,
		Usage: "update the language to be used in the hosted payment pages",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  RPMToCountryArg,
		Usage: "parameter to restrict the payment methods available to your customer to those from a single country",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:    SequenceTypeArg,
		Usage:   "indicate which type of payment this is in a recurring sequence",
		Default: string(mollie.OneOffSequence),
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  CustomerIDArg,
		Usage: "the ID of the Customer for whom the payment is being created",
	})
	commander.AddFlag(cpp, commander.FlagConfig{
		Name:  MandateIDArg,
		Usage: "when creating recurring payments, the ID of a specific Mandate may be supplied",
	})
}

func createPaymentAction(cmd *cobra.Command, args []string) {
	amount := ParseStringFromFlags(cmd, AmountValueArg)
	currency := ParseStringFromFlags(cmd, AmountCurrencyArg)
	desc := ParseStringFromFlags(cmd, DescriptionArg)
	rURL := ParseStringFromFlags(cmd, RedirectURLArg)
	whURL := ParseStringFromFlags(cmd, WebhookURLArg)
	meta := ParseStringFromFlags(cmd, MetadataArg)
	method := ParseStringFromFlags(cmd, MethodArg)
	locale := ParseStringFromFlags(cmd, LocaleArg)
	rpmCountry := ParseStringFromFlags(cmd, RPMToCountryArg)
	sequence := ParseStringFromFlags(cmd, SequenceTypeArg)
	mandate := ParseStringFromFlags(cmd, MandateIDArg)
	customer := ParseStringFromFlags(cmd, CustomerIDArg)

	if verbose {
		logger.Infof("creating payment of %s %s", amount, currency)
		PrintNonemptyFlagValue(RedirectURLArg, rURL)
		PrintNonemptyFlagValue(DescriptionArg, desc)
		PrintNonemptyFlagValue(WebhookURLArg, whURL)
		PrintNonemptyFlagValue(MetadataArg, meta)
		PrintNonemptyFlagValue(MethodArg, method)
		PrintNonemptyFlagValue(LocaleArg, locale)
		PrintNonemptyFlagValue(RPMToCountryArg, rpmCountry)
		PrintNonemptyFlagValue(CustomerIDArg, customer)
		PrintNonemptyFlagValue(MandateIDArg, mandate)
		PrintNonemptyFlagValue(SequenceTypeArg, sequence)
	}

	l := mollie.Locale(locale)
	c := mollie.Locale(rpmCountry)
	s := mollie.SequenceType(sequence)
	m := mollie.PaymentMethod(method)

	p := mollie.Payment{
		Amount: &mollie.Amount{
			Currency: currency,
			Value:    amount,
		},
		Description:                     desc,
		RedirectURL:                     rURL,
		WebhookURL:                      whURL,
		Metadata:                        meta,
		Locale:                          l,
		RestrictPaymentMethodsToCountry: c,
		CustomerID:                      customer,
		MandateID:                       mandate,
		SequenceType:                    s,
		Method:                          m,
	}

	p, err := API.Payments.Create(p)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
		logger.Infof("payment successfully created")
		logger.Infof("Payment created at %s", p.CreatedAt)
	}

	if json {
		printJSONP(p)
	}

	disp := displayers.MolliePayment{Payment: &p}

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
