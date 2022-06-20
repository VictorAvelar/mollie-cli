package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listInvoicesCmd(p *commander.Command) *commander.Command {
	li := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			Example:   "mollie invoices list",
			ShortDesc: "Retrieve all invoices on the account. Optionally filter on year or invoice number.",
			Execute:   listInvoicesAction,
		},
		invoicesCols(),
	)

	commander.AddFlag(li, commander.FlagConfig{
		Name:  ReferenceArg,
		Usage: "ilter for an invoice with a specific invoice number / reference",
	})

	commander.AddFlag(li, commander.FlagConfig{
		Name:  YearArg,
		Usage: "ilter for invoices from a specific year (e.g. 2020)",
	})

	AddFromFlag(li)
	AddLimitFlag(li)

	return li
}

func listInvoicesAction(cmd *cobra.Command, args []string) {
	ref := ParseStringFromFlags(cmd, ReferenceArg)
	year := ParseStringFromFlags(cmd, YearArg)
	from := ParseStringFromFlags(cmd, FromArg)
	limit := ParseIntFromFlags(cmd, LimitArg)

	opts := &mollie.InvoicesListOptions{
		Reference: ref,
		Year:      year,
		From:      from,
		Limit:     int64(limit),
	}

	res, is, err := app.API.Invoices.List(context.Background(), opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Invoices, is, res)

	if verbose {
		app.Logger.Infof("retrieved %d invoices", is.Count)
		app.Logger.Infof("request target: %s", is.Links.Self.Href)
		app.Logger.Infof("request docs: %s", is.Links.Documentation.Href)
	}

	disp := displayers.MollieInvoiceList{
		InvoicesList: is,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Invoices),
			invoicesCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
