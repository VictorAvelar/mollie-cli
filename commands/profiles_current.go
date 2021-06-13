package commands

import (
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
			LongDesc: `Use this API if you are creating a plugin or SaaS application that allows users to enter a Mollie API key, 
and you want to give a confirmation of the website profile that will be used in your plugin 
or application.`,
		},
		getProfileCols(),
	)
}

func currentProfileAction(cmd *cobra.Command, args []string) {
	p, err := API.Profiles.Current()
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
	}

	disp := displayers.MollieProfile{Profile: p}

	err = printer.Display(
		&disp,
		display.FilterColumns(
			parseFieldsFromFlag(cmd),
			getProfileCols(),
		),
	)
	if err != nil {
		logger.Error(err)
	}
}
