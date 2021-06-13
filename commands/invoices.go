package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func invoices() *commander.Command {
	i := commander.Builder(
		nil,
		commander.Config{
			Namespace: "invoices",
			Example:   "mollie invoices",
			ShortDesc: "Operations over Mollie's Invoices API.",
		},
		invoicesCols(),
	)

	listInvoicesCmd(i)
	getInvoicesCmd(i)

	return i
}

func invoicesCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"REFERENCE",
		"VAT_NUMBER",
		"STATUS",
		"ISSUED_AT",
		"PAID_AT",
		"DUE_AT",
		"NET_AMOUNT",
		"VAT_AMOUNT",
		"GROSS_AMOUNT",
	}
}
