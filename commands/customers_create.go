package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func createCustomerCmd(p *commander.Command) *commander.Command {
	cc := commander.Builder(
		p,
		commander.Config{
			Namespace: "create",
			Aliases:   []string{"new", "add"},
			ShortDesc: "Creates a simple minimal representation of a customer.",
			Example:   "mollie customers create --name 'test customer' --email test@example.com",
			Execute:   createCustomerAction,
			PostHook:  printJSONAction,
		},
		customersCols(),
	)

	commander.AddFlag(cc, commander.FlagConfig{
		Name:  NameArg,
		Usage: "the full name of the customer",
	})

	commander.AddFlag(cc, commander.FlagConfig{
		Name:  EmailArg,
		Usage: "the email address of the customer",
	})

	AddLocaleFlag(cc)

	commander.AddFlag(cc, commander.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like, and we will save the data alongside the customer",
	})

	return cc
}

func createCustomerAction(cmd *cobra.Command, args []string) {
	var c mollie.Customer
	{
		name := ParseStringFromFlags(cmd, NameArg)
		email := ParseStringFromFlags(cmd, EmailArg)
		locale := ParseStringFromFlags(cmd, LocaleArg)
		meta := ParseStringFromFlags(cmd, MetadataArg)

		c = mollie.Customer{
			Email:    email,
			Name:     name,
			Locale:   mollie.Locale(locale),
			Metadata: map[string]interface{}{"metadata": meta},
		}
	}

	res, nc, err := app.API.Customers.Create(context.Background(), c)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Customers, nc, res)

	if verbose {
		app.Logger.Infof("request target: %s", nc.Links.Self.Href)
		app.Logger.Infof("request docs: %s", nc.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomer{
		Customer: nc,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Customers),
			customersCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
