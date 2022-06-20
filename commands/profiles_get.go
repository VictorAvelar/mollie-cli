package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getProfileCmd(p *commander.Command) *commander.Command {
	gp := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve details of a profile, using the profile’s identifier.",
			Execute:   getProfileAction,
			Example:   "mollie profiles get --id=pfl_token",
			PostHook:  printJsonAction,
		},
		getProfileCols(),
	)

	AddIDFlag(gp, true)

	return gp
}

func getProfileAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	res, p, err := app.API.Profiles.Get(context.Background(), id)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Profiles, p, res)

	if verbose {
		app.Logger.Infof("using profile id: %s", id)
		app.Logger.Infof("request target: %s", p.Links.Self.Href)
	}

	disp := displayers.MollieProfile{Profile: p}

	err = app.Printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd, Profiles),
			getProfileCols(),
		),
	)
	if err != nil {
		app.Logger.Error(err)
	}
}
