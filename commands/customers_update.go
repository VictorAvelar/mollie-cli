package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func updateCustomerCmd(p *commander.Command) *commander.Command {
	uc := commander.Builder(
		p,
		commander.Config{
			Namespace: "update",
			Aliases:   []string{"edit", "change", "mutate"},
			ShortDesc: "Updates an existing customer.",
			Example:   "mollie customers update --name 'new name'",
			Execute:   updateCustomerAction,
			PostHook:  printJsonAction,
		},
		customersCols(),
	)

	AddIDFlag(uc, true)

	commander.AddFlag(uc, commander.FlagConfig{
		Name:  NameArg,
		Usage: "the full name of the customer",
	})

	commander.AddFlag(uc, commander.FlagConfig{
		Name:  EmailArg,
		Usage: "the email address of the customer",
	})

	AddLocaleFlag(uc)

	commander.AddFlag(uc, commander.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like, and we will save the data alongside the customer",
	})

	return uc
}

func updateCustomerAction(cmd *cobra.Command, args []string) {
	var c mollie.Customer
	{
		id := ParseStringFromFlags(cmd, IDArg)
		name := ParseStringFromFlags(cmd, NameArg)
		email := ParseStringFromFlags(cmd, EmailArg)
		locale := ParseStringFromFlags(cmd, LocaleArg)
		meta := ParseStringFromFlags(cmd, LocaleArg)

		c = mollie.Customer{
			ID:       id,
			Email:    email,
			Name:     name,
			Locale:   mollie.Locale(locale),
			Metadata: map[string]interface{}{"metadata": meta},
		}
	}

	res, uc, err := app.API.Customers.Update(context.Background(), c.ID, c)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Customers, uc, res)

	if verbose {
		app.Logger.Infof("request target: %s", uc.Links.Self.Href)
		app.Logger.Infof("request docs: %s", uc.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomer{
		Customer: uc,
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
