package commands

import (
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
