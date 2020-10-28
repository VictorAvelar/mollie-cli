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
	p := &command.Command{
		Command: &cobra.Command{
			Use:     "payments",
			Short:   "All operations to handle payments",
			Aliases: []string{"pay", "p"},
		},
	}

	lp := command.Builder(
		p,
		"list",
		"Retrieve all payments created",
		`Retrieve all payments created with the current website profile, 
ordered from newest to oldest. The results are paginated.`,
		RunListPayments,
		paymentsCols,
	)

	command.AddStringFlag(lp, FromArg, "", "", "offset the result set to the payment with this ID.", false)
	command.AddIntFlag(lp, LimitArg, "", "the number of payments to return", 250, false)

	gp := command.Builder(
		p,
		"get",
		"Retrieve a single payment object by its payment token.",
		``,
		RunGetPayment,
		[]string{},
	)

	command.AddStringFlag(gp, IDArg, "", "", "the payment token/id", true)

	cp := command.Builder(
		p,
		"cancel",
		"Cancel a payment by its payment token.",
		``,
		RunCancelPayment,
		[]string{},
	)

	command.AddStringFlag(cp, IDArg, "", "", "the payment token/id", true)

	cpp := command.Builder(
		p,
		"create",
		"Payment creation",
		``,
		RunCreatePayment,
		[]string{},
	)

	// Add common aliases
	cpp.Aliases = []string{"new", "start"}
	command.AddStringFlag(cpp, AmountValueArg, "", "", "a string containing the exact amount you want to charge in the given currency", true)
	command.AddStringFlag(cpp, AmountCurrencyArg, "", "", "an ISO 4217 currency code", true)
	command.AddStringFlag(cpp, DescriptionArg, "", "", "the description of the payment you’re creating to be show to your customers when possible", true)
	command.AddStringFlag(cpp, RedirectURLArg, "", "", "the URL your customer will be redirected to after the payment process", true)
	command.AddStringFlag(cpp, WebhookURLArg, "", "", "set the webhook URL, where we will send payment status updates to", false)
	command.AddStringFlag(cpp, MetadataArg, "", "", "any data you like, for example a string or a JSON object", false)
	command.AddStringFlag(cpp, MethodArg, "", "", "change the payment to a different payment method.", false)
	command.AddStringFlag(cpp, LocaleArg, "", "", "update the language to be used in the hosted payment pages", false)
	command.AddStringFlag(cpp, RPMToCountryArg, "", "", "parameter to restrict the payment methods available to your customer to those from a single country", false)
	command.AddStringFlag(cpp, SequenceTypeArg, "", string(mollie.OneOffSequence), "indicate which type of payment this is in a recurring sequence", false)
	command.AddStringFlag(cpp, CustomerIDArg, "", "", "the ID of the Customer for whom the payment is being created", false)
	command.AddStringFlag(cpp, MandateIDArg, "", "", "when creating recurring payments, the ID of a specific Mandate may be supplied", false)

	up := command.Builder(
		p,
		"update",
		"Update some details of a created payment",
		`There are also payment method specific parameters available, check the docs, please.`,
		RunUpdatePayment,
		[]string{},
	)

	command.AddStringFlag(up, IDArg, "", "", "the payment token/id", true)
	command.AddStringFlag(up, DescriptionArg, "", "", "the description of the payment.", false)
	command.AddStringFlag(up, RedirectURLArg, "", "", "the URL your customer will be redirected to after the payment process.", false)
	command.AddStringFlag(up, WebhookURLArg, "", "", "set the webhook URL, where we will send payment status updates to.", false)
	command.AddStringFlag(up, MetadataArg, "", "", "any data you like, for example a string or a JSON object", false)
	command.AddStringFlag(up, MethodArg, "", "", "change the payment to a different payment method.", false)
	command.AddStringFlag(up, LocaleArg, "", "", "update the language to be used in the hosted payment pages", false)
	command.AddStringFlag(up, RPMToCountryArg, "", "", "parameter to restrict the payment methods available to your customer to those from a single country", false)

	return p
}

// RunListPayments retrieves a list of payments for the current
// profile.
func RunListPayments(cmd *cobra.Command, args []string) {
	ps, err := API.Payments.List(nil)
	if err != nil {
		logger.Fatal(err)
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

	if Verbose {
		logger.Infof("retrieving payment with id (token) %s", id)
	}

	p, err := API.Payments.Get(id, nil)
	if err != nil {
		logger.Fatal(err)
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

	if Verbose {
		logger.Infof("canceling payment with id (token) %s", id)
	}

	p, err := API.Payments.Cancel(id)
	if err != nil {
		logger.Fatal(err)
	}

	if Verbose {
		logger.Info("payment successfully cancelled")
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

	if Verbose {
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

	if Verbose {
		logger.Info("payment successfully created")
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

	if Verbose {
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

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logger.Fatal(err)
	}
}
