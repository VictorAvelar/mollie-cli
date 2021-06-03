package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

var (
	methodsCols = []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"ISSUERS",
		"MIN_AMOUNT",
		"MAX_AMOUNT",
		"LOGO",
	}
)

func methods() *commander.Command {
	m := commander.Builder(nil, commander.Config{
		Namespace: "methods",
		Aliases:   []string{"vendors", "meths"},
		ShortDesc: "All payment methods that Mollie offers and can be activated",
	}, methodsCols)

	// Add namespace persistent flags.
	AddIncludeFlag(m, true)
	AddPrompterFlag(m, true)

	// Add child commands.
	listPaymentMethodsCmd(m)
	allPaymentMethodsCmd(m)
	getPaymentMethodCmd(m)

	return m
}
