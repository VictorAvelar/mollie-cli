package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
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

	res, ps, err := app.API.Payments.List(context.Background(), &opts)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Payments, ps, res)

	if verbose {
		app.Logger.Infof("retrieved %d payments", ps.Count)
		app.Logger.Infof("request target: %s", ps.Links.Self.Href)
		app.Logger.Infof("request docs: %s", ps.Links.Documentation.Href)
	}

	disp := displayers.MollieListPayments{
		PaymentList: ps,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Payments),
			getPaymentCols(),
		),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
