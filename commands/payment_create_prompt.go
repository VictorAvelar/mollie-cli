package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
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

	if verbose {
		logger.Info("Collected payment from prompter: %v", payment)
	}

	res, err := API.Payments.Create(payment)
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
		PrintJsonP(res)
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

func promptStringClean(q, def string) string {
	v, err := prompter.String(q, def)
	if err != nil {
		logger.Fatal(err)
	}

	return v
}
