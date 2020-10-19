package commands

import (
	"fmt"

	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	paymentsCols = []string{}
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

	return p
}

// RunListPayments retrieves a list of payments for the current
// profile.
func RunListPayments(cmd *cobra.Command, args []string) {
	ps, err := API.Payments.List(nil)
	if err != nil {
		logrus.Infof("%+v", err)
	}

	fmt.Printf("%+v\n", ps)
}
