package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func order() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "order",
			ShortDesc:          "The Orders API allows you to create payment intents with order management functionalities",
			PostHook:           printJsonAction,
			PersistentPostHook: printCurl,
		},
		getorderCols(),
	)

	getOrderCmd(p)

	return p
}

func getorderCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.order.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}
