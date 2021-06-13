package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listPaymentsCmd(p *commander.Command) *commander.Command {
	lp := commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all payments created",
			LongDesc: `Retrieve all payments created with the current website profile, 
ordered from newest to oldest. The results are paginated.`,
			Execute: listPaymentsAction,
			Example: "mollie payments list --limit=3",
		},
		getPaymentCols(),
	)

	AddFromFlag(lp)
	AddLimitFlag(lp)
	AddEmbedFlag(lp)
	AddIncludeFlag(lp, false)

	return lp
}

func listPaymentsAction(cmd *cobra.Command, args []string) {
	var opts mollie.ListPaymentOptions
	{
		opts.Limit = ParseIntFromFlags(cmd, LimitArg)
		opts.From = ParseStringFromFlags(cmd, FromArg)
		opts.Embed = ParseStringFromFlags(cmd, EmbedArg)
	}
	ps, err := API.Payments.List(&opts)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("retrieved %d payments", ps.Count)
		logger.Infof("request target: %s", ps.Links.Self.Href)
		logger.Infof("request docs: %s", ps.Links.Documentation.Href)
	}

	if json {
		printJSONP(ps)
	}

	disp := displayers.MollieListPayments{
		PaymentList: &ps,
	}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			getPaymentCols(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
