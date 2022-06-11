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

		if verbose {
			PrintNonEmptyFlags(cmd)
		}

		c = mollie.Customer{
			ID:       id,
			Email:    email,
			Name:     name,
			Locale:   mollie.Locale(locale),
			Metadata: map[string]interface{}{"metadata": meta},
		}
	}

	_, uc, err := API.Customers.Update(context.Background(), c.ID, c)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", uc.Links.Self.Href)
		logger.Infof("request docs: %s", uc.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomer{
		Customer: uc,
	}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			customersCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
