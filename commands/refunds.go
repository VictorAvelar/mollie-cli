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
			Namespace: "refunds",
			Aliases:   []string{"refs", "rf"},
			ShortDesc: "All operations to handle refunds",
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
	return []string{
		"RESOURCE",
		"ID",
		"AMOUNT",
		"SETTLEMENT_ID",
		"SETTLEMENT_AMOUNT",
		"DESCRIPTION",
		"METADATA",
		"STATUS",
		"PAYMENT_ID",
		"ORDER_ID",
		"CREATED_AT",
	}
}

func getRefundList(opts *mollie.ListRefundOptions, payment string) (*mollie.RefundList, error) {
	if payment != "" {
		_, rl, err := API.Refunds.ListRefundPayment(
			context.Background(), payment, opts,
		)
		return rl, err
	}

	_, rl, err := API.Refunds.ListRefund(context.Background(), opts)

	return rl, err
}
