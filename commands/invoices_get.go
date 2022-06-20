package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getInvoicesCmd(p *commander.Command) *commander.Command {
	gi := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			Execute:   getInvoicesAction,
			Example:   "mollie invoices get --id inv_test",
			LongDesc: `Retrieve details of an invoice, using the invoice’s identifier.
If you want to retrieve the details of an invoice by its invoice number, 
use the list endpoint with the reference parameter.`,
		},
		invoicesCols(),
	)

	AddIDFlag(gi, true)

	return gi
}

func getInvoicesAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, i, err := app.API.Invoices.Get(context.Background(), id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Invoices, i, res)

	if verbose {
		app.Logger.Infof("request target: %s", i.Links.Self.Href)
		app.Logger.Infof("request docs: %s", i.Links.Documentation.Href)
	}

	disp := &displayers.MollieInvoice{
		Invoice: i,
	}

	err = app.Printer.Display(
		disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Invoices),
			invoicesCols(),
		),
	)

	if err != nil {
		app.Logger.Fatal(err)
	}
}
