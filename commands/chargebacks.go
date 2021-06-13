package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func chargebacks() *commander.Command {
	cb := commander.Builder(
		nil,
		commander.Config{
			Namespace: "chargebacks",
			ShortDesc: "Operations with the Chargebacks API",
			Aliases:   []string{"cb", "cback"},
		},
		getChargebacksCols(),
	)

	getChargebacksCmd(cb)
	listChargebacksCmd(cb)

	return cb
}

func getChargebacksCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"AMOUNT",
		"SETTLEMENT_AMOUNT",
		"CREATED_AT",
		"REVERSED_AT",
		"PAYMENT_ID",
	}
}
