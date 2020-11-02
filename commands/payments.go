package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	paymentsCols = []string{
		"ID",
		"Mode",
		"Created",
		"Expires",
		"Cancelable",
		"Amount",
		"Method",
		"Description",
	}
)

// Payments builds the payments command tree.
func Payments() *command.Command {
	p := command.Builder(
		nil,
		command.Config{
			Namespace: "payments",
			ShortDesc: "All operations to handle payments",
			Aliases:   []string{"pay", "p"},
		},
		noCols,
	)

	lp := command.Builder(
		p,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all payments created",
			LongDesc: `Retrieve all payments created with the current website profile, 
ordered from newest to oldest. The results are paginated.`,
			Execute: RunListPayments,
			Example: "mollie payments list --limit=3",
		},
		paymentsCols,
	)

	command.AddFlag(lp, command.FlagConfig{
		Name:  FromArg,
		Usage: "offset the result set to the payment with this ID.",
	})
	command.AddFlag(lp, command.FlagConfig{
		Name:  LimitArg,
		Usage: "the number of payments to return",
	})

	gp := command.Builder(
		p,
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single payment object by its payment token.",
			Execute:   RunGetPayment,
			Example:   "mollie payments get --id=tr_token",
		},
		noCols,
	)

	command.AddFlag(gp, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment token/id",
		Required: true,
	})

	cp := command.Builder(
		p,
		command.Config{
			Namespace: "cancel",
			ShortDesc: "Cancel a payment by its payment token.",
			Execute:   RunCancelPayment,
			Example:   "mollie payments cancel --id=tr_token",
		},
		noCols,
	)

	command.AddFlag(cp, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment token/id",
		Required: true,
	})

	cpp := command.Builder(
		p,
		command.Config{
			Namespace: "create",
			ShortDesc: "Create a new payment",
			Execute:   RunCreatePayment,
			Aliases:   []string{"new", "start"},
			Example:   "mollie payments create --amount-value=200.00 --amount-currency=USD --redirect-to=https://victoravelar.com --description='custom example payment'",
		},
		noCols,
	)

	command.AddFlag(cpp, command.FlagConfig{
		Name:     AmountValueArg,
		Usage:    "a string containing the exact amount you want to charge in the given currency",
		Required: true,
	})
	command.AddFlag(cpp, command.FlagConfig{
		FlagType: command.BoolFlag,
		Name:     CancelableArg,
		Usage:    "indicates if the payment can be cancelled",
		Default:  true,
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:     AmountCurrencyArg,
		Usage:    "an ISO 4217 currency code",
		Required: true,
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:     DescriptionArg,
		Usage:    "the description of the payment you’re creating to be show to your customers when possible",
		Required: true,
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:     RedirectURLArg,
		Usage:    "the URL your customer will be redirected to after the payment process",
		Required: true,
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  WebhookURLArg,
		Usage: "set the webhook URL, where we will send payment status updates to",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  MetadataArg,
		Usage: "any data you like, for example a string or a JSON object",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  MethodArg,
		Usage: "change the payment to a different payment method.",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "update the language to be used in the hosted payment pages",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  RPMToCountryArg,
		Usage: "parameter to restrict the payment methods available to your customer to those from a single country",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:    SequenceTypeArg,
		Usage:   "indicate which type of payment this is in a recurring sequence",
		Default: string(mollie.OneOffSequence),
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  CustomerIDArg,
		Usage: "the ID of the Customer for whom the payment is being created",
	})
	command.AddFlag(cpp, command.FlagConfig{
		Name:  MandateIDArg,
		Usage: "when creating recurring payments, the ID of a specific Mandate may be supplied",
	})

	up := command.Builder(
		p,
		command.Config{
			Namespace: "update",
			ShortDesc: "Update some details of a created payment",
			LongDesc:  `There are also payment method specific parameters available, check the docs, please.`,
			Execute:   RunUpdatePayment,
		},
		noCols,
	)

	command.AddFlag(up, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment token/id",
		Required: true,
	})
	command.AddFlag(up, command.FlagConfig{
		Name:     DescriptionArg,
		Usage:    "the description of the payment to show to your customers when possible",
		Required: true,
	})
	command.AddFlag(up, command.FlagConfig{
		Name:     RedirectURLArg,
		Usage:    "the URL your customer will be redirected to after the payment process",
		Required: true,
	})
	command.AddFlag(up, command.FlagConfig{
		Name:  WebhookURLArg,
		Usage: "set the webhook URL, where we will send payment status updates to",
	})
	command.AddFlag(up, command.FlagConfig{
		Name:  MetadataArg,
		Usage: "any data you like, for example a string or a JSON object",
	})
	command.AddFlag(up, command.FlagConfig{
		Name:  MethodArg,
		Usage: "change the payment to a different payment method.",
	})
	command.AddFlag(up, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "update the language to be used in the hosted payment pages",
	})
	command.AddFlag(up, command.FlagConfig{
		Name:  RPMToCountryArg,
		Usage: "parameter to restrict the payment methods available to your customer to those from a single country",
	})

	return p
}

// RunListPayments retrieves a list of payments for the current
// profile.
func RunListPayments(cmd *cobra.Command, args []string) {
	ps, err := API.Payments.List(nil)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("retrieved %d payments", ps.Count)
		logger.Infof("request target: %s", ps.Links.Self.Href)
		logger.Infof("request docs: %s", ps.Links.Docs.Href)
	}

	disp := displayers.MollieListPayments{
		PaymentList: &ps,
	}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetPayment retrieves a single payment object.
func RunGetPayment(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		logger.Infof("retrieving payment with id (token) %s", id)
	}

	p, err := API.Payments.Get(id, nil)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunCancelPayment cancels a payment.
func RunCancelPayment(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		logger.Infof("canceling payment with id (token) %s", id)
	}

	p, err := API.Payments.Cancel(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
		logger.Infof("payment successfully cancelled")
		logger.Infof("cancellation processed at %s", p.CanceledAt)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunCreatePayment instantiates a new payment using your mollie account.
func RunCreatePayment(cmd *cobra.Command, args []string) {
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
		Locale:                          &l,
		RestrictPaymentMethodsToCountry: &c,
		CustomerID:                      customer,
		MandateID:                       mandate,
		SequenceType:                    &s,
		Method:                          &m,
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

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}

// RunUpdatePayment mutates a payment method.
func RunUpdatePayment(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)
	desc := ParseStringFromFlags(cmd, DescriptionArg)
	rURL := ParseStringFromFlags(cmd, RedirectURLArg)
	whURL := ParseStringFromFlags(cmd, WebhookURLArg)
	meta := ParseStringFromFlags(cmd, MetadataArg)
	method := ParseStringFromFlags(cmd, MethodArg)
	locale := ParseStringFromFlags(cmd, LocaleArg)
	rpmCountry := ParseStringFromFlags(cmd, RPMToCountryArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
		PrintNonemptyFlagValue(RedirectURLArg, rURL)
		PrintNonemptyFlagValue(DescriptionArg, desc)
		PrintNonemptyFlagValue(WebhookURLArg, whURL)
		PrintNonemptyFlagValue(MetadataArg, meta)
		PrintNonemptyFlagValue(MethodArg, method)
		PrintNonemptyFlagValue(LocaleArg, locale)
		PrintNonemptyFlagValue(RPMToCountryArg, rpmCountry)
	}

	l := mollie.Locale(locale)
	c := mollie.Locale(rpmCountry)
	m := mollie.PaymentMethod(method)

	p, err := API.Payments.Update(id, mollie.Payment{
		RedirectURL:                     rURL,
		Description:                     desc,
		WebhookURL:                      whURL,
		Metadata:                        meta,
		Locale:                          &l,
		RestrictPaymentMethodsToCountry: &c,
		Method:                          &m,
	})
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
		logger.Infof("request docs: %s", p.Links.Documentation.Href)
		logger.Infof("payment successfully updated")
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}
