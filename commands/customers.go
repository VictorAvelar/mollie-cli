package commands

import "github.com/VictorAvelar/mollie-cli/internal/command"

var (
	customerCols = []string{
		"ID",
		"Mode",
		"Name",
		"Email",
		"Created At",
	}
)

// Customers creates the customers command tree.
func Customers() *command.Command {
	c := command.Builder(
		nil,
		command.Config{
			Namespace: "customers",
			ShortDesc: "Operations with customers API",
			Aliases:   []string{"cust", "cstm"},
		},
		noCols,
	)
	return c
}
