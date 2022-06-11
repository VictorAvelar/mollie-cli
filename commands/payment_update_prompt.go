package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func updatePaymentPromptCmd(p *commander.Command) *commander.Command {
	return commander.Builder(
		p,
		commander.Config{
			Namespace: "prompt",
			ShortDesc: "Updates a payment by prompting the user for information",
			Execute:   promptUpdatePaymentAction,
		},
		getPaymentCols(),
	)
}

func promptUpdatePaymentAction(cmd *cobra.Command, args []string) {
	var payment mollie.Payment
	{
		payment = mollie.Payment{}
		payment.ID = promptStringClean("payment id", "")
		payment.Description = promptStringClean("payment description", "payment from CLI")
		payment.RedirectURL = promptStringClean("redirect URL", "")
		payment.Method = promptPaymentMethod()
		payment.Locale = promptLocale()
		payment.Metadata = promptStringClean("custom metadata", "")
		payment.SequenceType = promptSequenceType()
		payment.WebhookURL = promptStringClean("webhook URL", "")
		payment.CustomerID = promptStringClean("customer id", "")
		attachPaymentMethodSpecificValues(&payment)
	}

	_, res, err := API.Payments.Update(
		context.Background(),
		payment.ID,
		payment,
	)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", res.Links.Self.Href)
		logger.Infof("request docs: %s", res.Links.Documentation.Href)
		logger.Infof("payment successfully created")
		logger.Infof("Payment created at %s", res.CreatedAt)
	}

	if json {
		printJSONP(res)
	}

	err = printer.Display(
		&displayers.MolliePayment{Payment: &payment},
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			getPaymentCols(),
		),
	)

	if err != nil {
		logger.Fatal(err)
	}

}
