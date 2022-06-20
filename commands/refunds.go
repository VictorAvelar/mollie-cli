package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/avocatl/admiral/pkg/commander"
)

func refunds() *commander.Command {
	r := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "refunds",
			Aliases:            []string{"refs", "rf"},
			ShortDesc:          "All operations to handle refunds",
			PostHook:           printJsonAction,
			PersistentPostHook: printCurl,
		},
		refundsCols(),
	)

	listRefundCmd(r)
	refundsGetCmd(r)
	cancelRefundCmd(r)
	allRefundsCmd(r)
	createRefundCmd(r)

	return r
}

func refundsCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.refunds.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}

func getRefundList(opts *mollie.ListRefundOptions, payment string) (*mollie.RefundList, error) {
	if payment != "" {
		_, rl, err := app.API.Refunds.ListRefundPayment(
			context.Background(), payment, opts,
		)
		return rl, err
	}

	res, rl, err := app.API.Refunds.ListRefund(context.Background(), opts)

	addStoreValues(Refunds, rl, res)

	return rl, err
}
