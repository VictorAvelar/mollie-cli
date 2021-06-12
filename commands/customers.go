package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func customers() *commander.Command {
	c := commander.Builder(
		nil,
		commander.Config{
			Namespace: "customers",
			ShortDesc: "Operations with customers API.",
			Aliases:   []string{"cust", "cstm"},
		},
		customersCols(),
	)

	getCustomersCmd(c)
	listCustomerCmd(c)
	createCustomerCmd(c)
	updateCustomerCmd(c)
	deleteCustomerCmd(c)

	return c
}

func customersCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"EMAIL",
		"LOCALE",
		"METADATA",
		"CREATED_AT",
	}
}
