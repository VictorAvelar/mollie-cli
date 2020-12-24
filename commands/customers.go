package commands

import (
	"fmt"
	"strconv"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	customerCols = []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"EMAIL",
		"LOCALE",
		"METADATA",
		"CREATED_AT",
	}
)

// Customers creates the customers command tree.
func Customers() *command.Command {
	c := command.Builder(
		nil,
		command.Config{
			Namespace: "customers",
			ShortDesc: "Operations with customers API.",
			Aliases:   []string{"cust", "cstm"},
		},
		noCols,
	)

	gc := command.Builder(
		c,
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single customer by its ID.",
			Example:   "mollie customers get --id=cs_token",
			Execute:   RunGetCustomer,
		},
		customerCols,
	)

	command.AddFlag(gc, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the customer id/token",
		Required: true,
	})

	lc := command.Builder(
		c,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieves all customers created.",
			Example:   "mollie customers list",
			Execute:   RunListCustomers,
		},
		customerCols,
	)

	command.AddFlag(lc, command.FlagConfig{
		Name:     LimitArg,
		FlagType: command.IntFlag,
		Usage:    "the number of customers to return (with a maximum of 250)",
		Default:  50,
	})

	command.AddFlag(lc, command.FlagConfig{
		Name:  FromArg,
		Usage: "offset the result set to the customer with this ID",
	})

	cc := command.Builder(
		c,
		command.Config{
			Namespace: "create",
			Aliases:   []string{"new", "add"},
			ShortDesc: "Creates a simple minimal representation of a customer.",
			Example:   "mollie customers create --name 'test customer' --email test@example.com",
			Execute:   RunCreateCustomer,
		},
		customerCols,
	)

	command.AddFlag(cc, command.FlagConfig{
		Name:  NameArg,
		Usage: "the full name of the customer",
	})

	command.AddFlag(cc, command.FlagConfig{
		Name:  EmailArg,
		Usage: "the email address of the customer",
	})

	command.AddFlag(cc, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "allows you to preset the language to be used in the hosted payment pages shown to the consumer",
	})

	command.AddFlag(cc, command.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like, and we will save the data alongside the customer",
	})

	uc := command.Builder(
		c,
		command.Config{
			Namespace: "update",
			Aliases:   []string{"edit", "change", "mutate"},
			ShortDesc: "Updates an existing customer.",
			Example:   "mollie customers update --name 'new name'",
			Execute:   RunUpdateCustomer,
		},
		customerCols,
	)

	command.AddFlag(uc, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the customer id/token",
		Required: true,
	})

	command.AddFlag(uc, command.FlagConfig{
		Name:  NameArg,
		Usage: "the full name of the customer",
	})

	command.AddFlag(uc, command.FlagConfig{
		Name:  EmailArg,
		Usage: "the email address of the customer",
	})

	command.AddFlag(uc, command.FlagConfig{
		Name:  LocaleArg,
		Usage: "allows you to preset the language to be used in the hosted payment pages shown to the consumer",
	})

	command.AddFlag(uc, command.FlagConfig{
		Name:  MetadataArg,
		Usage: "provide any data you like, and we will save the data alongside the customer",
	})

	dc := command.Builder(
		c,
		command.Config{
			Namespace: "delete",
			Aliases:   []string{"remove", "del"},
			ShortDesc: "Deletes a customer by its ID.",
			LongDesc:  "Deletes a customer. WARNING! All mandates and subscriptions created for this customer will be canceled as well.",
			Example:   "mollie customers delete --id cs_test",
			Execute:   RunDeleteCustomer,
		},
		customerCols,
	)

	command.AddFlag(dc, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the customer id/token",
		Required: true,
	})

	return c
}

// RunGetCustomer retrieves a customer by its id.
func RunGetCustomer(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
	}

	c, err := API.Customers.Get(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", c.Links.Self.Href)
		logger.Infof("request docs: %s", c.Links.Documentation.Href)
	}

	fmt.Printf("%+v", c)
}

// RunListCustomers retrieves all the created customers for the account.
func RunListCustomers(cmd *cobra.Command, args []string) {
	var opts *mollie.ListCustomersOptions
	{
		opts.Limit = ParseIntFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
	}

	if verbose {
		PrintNonemptyFlagValue(LimitArg, strconv.Itoa(opts.Limit))
		PrintNonemptyFlagValue(FromArg, opts.From)
	}

	cl, err := API.Customers.List(opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", cl.Links.Self.Href)
		logger.Infof("request docs: %s", cl.Links.Documentation.Href)
	}

	fmt.Printf("%+v", cl)
}

// RunCreateCustomer creates a simple minimal representation of a customer.
func RunCreateCustomer(cmd *cobra.Command, args []string) {
	var c mollie.Customer
	{
		name := ParseStringFromFlags(cmd, NameArg)
		email := ParseStringFromFlags(cmd, EmailArg)
		locale := ParseStringFromFlags(cmd, LocaleArg)
		meta := ParseStringFromFlags(cmd, LocaleArg)

		if verbose {
			PrintNonemptyFlagValue(NameArg, name)
			PrintNonemptyFlagValue(EmailArg, email)
			PrintNonemptyFlagValue(LocaleArg, locale)
			PrintNonemptyFlagValue(MetadataArg, meta)
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

	fmt.Printf("%+v", nc)
}

// RunUpdateCustomer updates an existing customer.
func RunUpdateCustomer(cmd *cobra.Command, args []string) {
	var c mollie.Customer
	{
		id := ParseStringFromFlags(cmd, IDArg)
		name := ParseStringFromFlags(cmd, NameArg)
		email := ParseStringFromFlags(cmd, EmailArg)
		locale := ParseStringFromFlags(cmd, LocaleArg)
		meta := ParseStringFromFlags(cmd, LocaleArg)

		if verbose {
			PrintNonemptyFlagValue(IDArg, id)
			PrintNonemptyFlagValue(NameArg, name)
			PrintNonemptyFlagValue(EmailArg, email)
			PrintNonemptyFlagValue(LocaleArg, locale)
			PrintNonemptyFlagValue(MetadataArg, meta)
		}

		c = mollie.Customer{
			ID:       id,
			Email:    email,
			Name:     name,
			Locale:   mollie.Locale(locale),
			Metadata: map[string]interface{}{"metadata": meta},
		}
	}

	uc, err := API.Customers.Update(c.ID, c)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", uc.Links.Self.Href)
		logger.Infof("request docs: %s", uc.Links.Documentation.Href)
	}

	fmt.Printf("%+v", uc)
}

// RunDeleteCustomer removes a customer by its id/token.
func RunDeleteCustomer(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(IDArg, id)
	}

	err := API.Customers.Delete(id)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("removed customer with id/token: %s", id)
	}

	fmt.Printf("%+v", id)
}
