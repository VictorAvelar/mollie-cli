package commands

import (
	"context"

	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func cancelRefundCmd(p *commander.Command) *commander.Command {
	dr := commander.Builder(
		p,
		commander.Config{
			Namespace: "cancel",
			Aliases:   []string{"delete", "remove", "cncl"},
			ShortDesc: "for certain payment methods where cancelation is possible.",
			LongDesc: `For certain payment methods, like iDEAL, the underlying banking system will delay refunds
until the next day. Until that time, refunds may be canceled manually in the Mollie Dashboard, 
or programmatically by using this endpoint.

A Refund can only be canceled while its status field is either queued or pending.`,
			Example:  "mollie refunds cancel --id=rf_test --payment=tr_test",
			Execute:  cancelRefundAction,
			PostHook: printJsonAction,
		},
		refundsCols(),
	)

	AddIDFlag(dr, true)
	AddPaymentFlag(dr)

	return dr
}

func cancelRefundAction(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	id := ParseStringFromFlags(cmd, IDArg)

	_, err := app.API.Refunds.Cancel(context.Background(), payment, id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	display.Text("*", "Refund successfully cancelled")
}
