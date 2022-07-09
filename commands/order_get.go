package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getOrderCmd(p *commander.Command) *commander.Command {
	gor := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single order by its ID",
			Execute:   getOrderAction,
			Example:   "mollie order get --id=pfl_token",
			PostHook:  printJsonAction,
		},
		getorderCols(),
	)

	AddIDFlag(gor, true)

	return gor
}

func getOrderAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, p, err := app.API.Orders.Get(context.Background(), id, &mollie.OrderOptions{})
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Order, p, res)

	if verbose {
		app.Logger.Infof("using order id: %s", id)
		app.Logger.Infof("request target: %s", p.Links.Self.Href)
	}

	disp := displayers.MollieOrder{
		Order: p,
	}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Order),
			getorderCols(),
		),
	)
	if err != nil {
		app.Logger.Error(err)
	}
}
