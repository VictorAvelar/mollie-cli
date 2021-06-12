package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func captures() *commander.Command {
	c := commander.Builder(
		nil,
		commander.Config{
			Namespace: "captures",
			ShortDesc: "Operations with Captures API.",
		},
		getCapturesCols(),
	)

	listCapturesCmd(c)
	getCapturesCmd(c)

	return c
}

func getCapturesCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"AMOUNT",
		"SETTLEMENT_AMOUNT",
		"PAYMENT_ID",
		"SHIPMENT_ID",
		"SETTLEMENT_ID",
		"CREATED_AT",
	}
}
