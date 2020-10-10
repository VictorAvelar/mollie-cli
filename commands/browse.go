package commands

import (
	"fmt"
	"os"

	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/VictorAvelar/mollie-cli/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	urlMap = map[string]map[string]string{
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
)

// Browse creates the browse commands tree
func Browse() *command.Command {
	b := &command.Command{
		Command: &cobra.Command{
			Use:     "browse",
			Args:    validateBrowseArgs,
			Example: "mollie browse web help-desk",
			Short:   "Browse through mollie's web resources",
			Long: `This command will open a browser tab with the target url for the required resource.
			Resources are grouped by category, and valid categories are:
			- auth
			- dashboard
			- developers
			- github
			- settings
			- web
			Use: mollie open -i for listing the complete resource tree.`,
			Run: RunOpenAction,
		},
	}

	command.AddBoolFlag(b, "info", "i", "prints extended info about the available resources", false, false)

	return b
}

func openNames() []string {
	names := make([]string, 0, len(urlMap))
	for k, v := range urlMap {
		names = append(names, fmt.Sprintf("--> %s", k))
		for x := range v {
			names = append(names, fmt.Sprintf("\t=> %s", x))
		}
	}
	return names
}

// RunOpenAction executes the open command
func RunOpenAction(cmd *cobra.Command, args []string) {
	in, err := cmd.Flags().GetBool("info")
	if err != nil {
		logrus.Fatal(err)
	}

	if in {
		for _, n := range openNames() {
			fmt.Println(n)
		}
		os.Exit(0)
	}
	var cat, page string
	{
		cat = args[0]
		page = args[1]
	}

	utils.Browse(urlMap[cat][page])
}

func validateBrowseArgs(cmd *cobra.Command, args []string) error {
	in, err := cmd.Flags().GetBool("info")
	if err != nil {
		logrus.Fatal(err)
	}

	if len(args) == 0 && in {
		return nil
	}

	var cat, page string
	{
		cat = args[0]
		page = args[1]
	}

	if _, ok := urlMap[cat]; !ok {
		return fmt.Errorf("invalid category name %s", cat)
	}

	if _, ok := urlMap[cat][page]; !ok {
		return fmt.Errorf("invalid page name %s", page)
	}

	return nil
}
