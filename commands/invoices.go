package commands

import (
	"strconv"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	invoicesCols = []string{
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
)

// Invoices creates the invoices command tree.
func Invoices() *command.Command {
	i := command.Builder(
		nil,
		command.Config{
			Namespace: "invoices",
			Example:   "mollie invoices",
			ShortDesc: "Operations over Mollie's Invoices API.",
		},
		noCols,
	)

	li := command.Builder(
		i,
		command.Config{
			Namespace: "list",
			Example:   "mollie invoices list",
			ShortDesc: "Retrieve all invoices on the account. Optionally filter on year or invoice number.",
			Execute:   RunListInvoices,
		},
		invoicesCols,
	)

	command.AddFlag(li, command.FlagConfig{
		Name:  ReferenceArg,
		Usage: "ilter for an invoice with a specific invoice number / reference",
	})

	command.AddFlag(li, command.FlagConfig{
		Name:  YearArg,
		Usage: "ilter for invoices from a specific year (e.g. 2020)",
	})

	command.AddFlag(li, command.FlagConfig{
		Name:  FromArg,
		Usage: "offset the result set to the invoice with this ID (pagination)",
	})

	command.AddFlag(li, command.FlagConfig{
		Name:     LimitArg,
		FlagType: command.IntFlag,
		Usage:    "the number of invoices to return (with a maximum of 250)",
		Default:  50,
	})

	gi := command.Builder(
		i,
		command.Config{
			Namespace: "get",
			Execute:   RunGetInvoice,
			Example:   "mollie invoices get --id inv_test",
			LongDesc: `Retrieve details of an invoice, using the invoice’s identifier.
If you want to retrieve the details of an invoice by its invoice number, 
use the list endpoint with the reference parameter.`,
		},
		invoicesCols,
	)

	command.AddFlag(gi, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the invoice identifier/token.",
		Required: true,
	})

	return i
}

// RunListInvoices returns all invocies for the profile/account.
func RunListInvoices(cmd *cobra.Command, args []string) {
	ref := ParseStringFromFlags(cmd, ReferenceArg)
	year := ParseStringFromFlags(cmd, YearArg)
	from := ParseStringFromFlags(cmd, FromArg)
	limit := ParseIntFromFlags(cmd, LimitArg)

	if verbose {
		PrintNonemptyFlagValue(ReferenceArg, ref)
		PrintNonemptyFlagValue(YearArg, year)
		PrintNonemptyFlagValue(FromArg, from)
		v := strconv.Itoa(limit)
		PrintNonemptyFlagValue(LimitArg, v)
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

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), invoicesCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetInvoice returns an invoice by its identifier/token.
func RunGetInvoice(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
	}

	i, err := API.Invoices.Get(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", i.Links.Self.Href)
		logger.Infof("request docs: %s", i.Links.Documentation.Href)
	}

	disp := displayers.MollieInvoice{
		Invoice: &i,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), invoicesCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
