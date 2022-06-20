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
			PostHook:  printJsonAction,
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

	res, p, err := app.API.Payments.Update(
		context.Background(),
		payment.ID,
		payment,
	)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Payments, p, res)

	if verbose {
		app.Logger.Infof("request target: %s", p.Links.Self.Href)
		app.Logger.Infof("request docs: %s", p.Links.Documentation.Href)
		app.Logger.Infof("payment successfully created")
		app.Logger.Infof("Payment created at %s", p.CreatedAt)
	}

	err = app.Printer.Display(
		&displayers.MolliePayment{Payment: p},
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Payments),
			getPaymentCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}

}
