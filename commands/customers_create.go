package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/avocatl/admiral/pkg/commander"
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

		if verbose {
			PrintNonEmptyFlags(cmd)
		}

		c = mollie.Customer{
			Email:    email,
			Name:     name,
			Locale:   mollie.Locale(locale),
			Metadata: map[string]interface{}{"metadata": meta},
		}
	}

	nc, err := API.Customers.Create(c)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", nc.Links.Self.Href)
		logger.Infof("request docs: %s", nc.Links.Documentation.Href)
	}

	disp := displayers.MollieCustomer{
		Customer: nc,
	}

	err = printer.Display(
		&disp,
		command.FilterColumns(
			parseFieldsFromFlag(cmd),
			customersCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
