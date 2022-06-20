package commands

import (
	"context"

	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func deleteCustomerCmd(p *commander.Command) *commander.Command {
	dc := commander.Builder(
		p,
		commander.Config{
			Namespace: "delete",
			Aliases:   []string{"remove", "del"},
			ShortDesc: "Deletes a customer by its ID.",
			LongDesc:  "Deletes a customer. WARNING! All mandates and subscriptions created for this customer will be canceled as well.",
			Example:   "mollie customers delete --id cs_test",
			Execute:   deleteCustomerAction,
		},
		customersCols(),
	)

	AddIDFlag(dc, true)

	return dc
}

func deleteCustomerAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, err := app.API.Customers.Delete(context.Background(), id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Customers, id, res)

	if verbose {
		app.Logger.Infof("removed customer with id/token: %s", id)
	}

	display.Text("*", "Customer deleted successfully")
}
