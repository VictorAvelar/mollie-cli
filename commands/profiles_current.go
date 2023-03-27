package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func currentProfileCmd(p *commander.Command) *commander.Command {
	return commander.Builder(
		p,
		commander.Config{
			Namespace: "current",
			Execute:   currentProfileAction,
			ShortDesc: "Retrieve details of the profile associated to the current API token.",
			Example:   "mollie profiles current",
			LongDesc: `
Use this API if you are creating a plugin or SaaS application that allows users to enter a Mollie API key,
and you want to give a confirmation of the website profile that will be used in your plugin
or application.`,
			PostHook: printJsonAction,
		},
		getProfileCols(),
	)
}

func currentProfileAction(cmd *cobra.Command, args []string) {
	res, p, err := app.API.Profiles.Current(context.Background())
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Profiles, p, res)

	if verbose {
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
