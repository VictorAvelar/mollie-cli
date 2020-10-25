package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/sirupsen/logrus"
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
		"retrieve all payments created",
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

	return p
}

// RunListPayments retrieves a list of payments for the current
// profile.
func RunListPayments(cmd *cobra.Command, args []string) {
	ps, err := API.Payments.List(nil)
	if err != nil {
		logrus.Fatal(err)
	}

	disp := displayers.MollieListPayments{
		PaymentList: &ps,
	}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}

// RunGetPayment retrieves a single payment object.
func RunGetPayment(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if Verbose {
		logrus.Infof("retrieving payment with id (token) %d", id)
	}

	p, err := API.Payments.Get(id, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}

// RunCancelPayment cancels a payment.
func RunCancelPayment(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if Verbose {
		logrus.Infof("canceling payment with id (token) %d", id)
	}

	p, err := API.Payments.Cancel(id)
	if err != nil {
		logrus.Fatal(err)
	}

	if Verbose {
		logrus.Info("payment successfully cancelled")
		logrus.Infof("cancellation processed at %s", p.CanceledAt)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}

// RunCreatePayment instantiates a new payment using your mollie account.
func RunCreatePayment(cmd *cobra.Command, args []string) {
	amount := ParseStringFromFlags(cmd, AmountValueArg)
	currency := ParseStringFromFlags(cmd, AmountCurrencyArg)
	desc := ParseStringFromFlags(cmd, DescriptionArg)
	rURL := ParseStringFromFlags(cmd, RedirectURLArg)

	if Verbose {
		logrus.Infof("creating payment of %s %s", amount, currency)
		logrus.Infof("redirect url received %s", rURL)
		logrus.Infof(`
this description will be shown to your customer in their 
payment provider statement or applications:
%s`, desc)
	}

	p := mollie.Payment{
		Amount: &mollie.Amount{
			Currency: currency,
			Value:    amount,
		},
		Description: desc,
		RedirectURL: rURL,
	}

	p, err := API.Payments.Create(p)
	if err != nil {
		logrus.Fatal(err)
	}

	if Verbose {
		logrus.Info("payment successfully created")
		logrus.Infof("Payment processed at %s", p.CreatedAt)
	}

	disp := displayers.MolliePayment{Payment: &p}

	err = command.Display(paymentsCols, disp.KV())
	if err != nil {
		logrus.Fatal(err)
	}
}
