package commands

import (
	"os"

	"github.com/VictorAvelar/mollie-cli/pkg/utils"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func browse() *commander.Command {
	b := commander.Builder(
		nil,
		commander.Config{
			Namespace: "browse [category] [resource]",
			Aliases:   []string{"resources", "open"},
			PreHook:   preBrowseHook,
			Execute:   browseAction,
			ShortDesc: "Browse through mollie's web resources",
			LongDesc: `This command will open a browser tab with the target url for the required resource.
Resources are grouped by category, and valid categories are:
- auth
- dashboard
- developers
- github
- settings
- web
Use: mollie open -i for listing the complete resource tree.`,
		},
		commander.NoCols(),
	)

	commander.AddFlag(b, commander.FlagConfig{
		Name:     "info",
		FlagType: commander.BoolFlag,
		Default:  false,
		Usage:    "prints the resources available to the output writer",
	})

	return b
}

func preBrowseHook(cmd *cobra.Command, args []string) {
	b, err := cmd.Flags().GetBool("info")
	if err != nil {
		logger.Fatal(err)
	}

	if b {
		printJSONP(getURLMap())
	}
}

func browseAction(cmd *cobra.Command, args []string) {
	resources := getURLMap()

	if len(args) != len("ok") {
		display.Text("X", "invalid number of arguments provided")
		os.Exit(1)
	}

	cat, open := args[0], args[1]

	utils.Browse(resources[cat][open])
}

func getURLMap() map[string]map[string]string {
	return map[string]map[string]string{
		"web": {
			"home":       "https://www.mollie.com",
			"developers": "https://www.mollie.com/en/developers",
			"packages":   "https://www.mollie.com/en/developers/packages",
			"help-desk":  "https://help.mollie.com/hc/",
		},
		"developers": {
			"guides":        "https://docs.mollie.com/index",
			"changelog":     "https://docs.mollie.com/changelog/v2/changelog",
			"apiref":        "https://docs.mollie.com/reference/v2/payments-api/create-payment",
			"payment-links": "https://useplink.com/en/",
		},
		"dashboard": {
			"onboarding":     "https://www.mollie.com/dashboard/onboarding",
			"payments":       "https://www.mollie.com/dashboard/payments",
			"refunds":        "https://www.mollie.com/dashboard/refunds",
			"chargebacks":    "https://www.mollie.com/dashboard/chargebacks",
			"orders":         "https://www.mollie.com/dashboard/orders",
			"administration": "https://www.mollie.com/dashboard/administration",
			"invoices":       "https://www.mollie.com/dashboard/invoices",
			"settlements":    "https://www.mollie.com/dashboard/settlements",
			"notifications":  "https://www.mollie.com/dashboard/notifications",
			"profiles":       "https://www.mollie.com/dashboard/profiles",
		},
		"auth": {
			"api-keys":   "https://www.mollie.com/dashboard/developers/api-keys",
			"org-tokens": "https://www.mollie.com/dashboard/developers/organization-access-tokens",
			"apps":       "https://www.mollie.com/dashboard/developers/applications",
		},
		"settings": {
			"organization":  "https://www.mollie.com/dashboard/organization",
			"team":          "https://www.mollie.com/dashboard/team",
			"bank-accounts": "https://www.mollie.com/dashboard/bank-accounts",
			"payouts":       "https://www.mollie.com/dashboard/payouts",
		},
		"github": {
			"source": "https://github.com/VictorAvelar/mollie-cli",
			"issues": "https://github.com/VictorAvelar/mollie-cli/issues",
		},
	}
}
