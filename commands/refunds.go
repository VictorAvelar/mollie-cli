package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
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
		return API.Refunds.ListRefundPayment(payment, opts)
	}

	return API.Refunds.ListRefund(opts)
}
