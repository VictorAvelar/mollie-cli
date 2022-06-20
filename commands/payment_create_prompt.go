package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/avocatl/admiral/pkg/prompter"
	"github.com/spf13/cobra"
)

func promptPaymentCmd(p *commander.Command) *commander.Command {
	return commander.Builder(
		p,
		commander.Config{
			Namespace: "prompt",
			ShortDesc: "Creates a payment by prompting the user for values",
			Execute:   promptPaymentAction,
			PostHook:  printJsonAction,
		},
		getPaymentCols(),
	)
}

func promptPaymentAction(cmd *cobra.Command, args []string) {
	var payment mollie.Payment
	{
		payment = mollie.Payment{}
		payment.Amount = promptAmount()
		payment.Description = promptStringClean("payment description", "payment from CLI")
		payment.RedirectURL = promptStringClean("redirect URL", "")
		payment.Method = promptPaymentMethod()
		payment.Locale = promptLocale()
		payment.Metadata = promptStringClean("custom metadata", "")
		payment.SequenceType = promptSequenceType()
		payment.WebhookURL = promptStringClean("webhook URL", "")
		payment.CustomerID = promptStringClean("customer id", "")
		attachPaymentMethodSpecificValues(&payment)
		attachAccessTokenParams(&payment)
	}

	res, p, err := app.API.Payments.Create(context.Background(), payment, nil)
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

func promptStringClean(q, def string) string {
	v, err := prompter.String(q, def)
	if err != nil {
		app.Logger.Fatal(err)
	}

	return v
}
