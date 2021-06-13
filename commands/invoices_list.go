package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/avocatl/admiral/pkg/commander"
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

	if verbose {
		PrintNonEmptyFlags(cmd)
	}

	opts := &mollie.ListInvoiceOptions{
		Reference: ref,
		Year:      year,
		From:      from,
		Limit:     int64(limit),
	}

	is, err := API.Invoices.List(opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("retrieved %d invoices", is.Count)
		logger.Infof("request target: %s", is.Links.Self.Href)
		logger.Infof("request docs: %s", is.Links.Documentation.Href)
	}

	disp := displayers.MollieInvoiceList{
		InvoiceList: &is,
	}

	err = printer.Display(
		&disp,
		command.FilterColumns(parseFieldsFromFlag(cmd), invoicesCols()),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
