package commands

import (
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
		},
		getProfileCols(),
	)

	AddIDFlag(gp, true)

	return gp
}

func getProfileAction(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		logger.Infof("fetching profile with id %s", id)
	}

	p, err := API.Profiles.Get(id)
	if err != nil {
		logger.Fatal(err)
	}
	if verbose {
		logger.Infof("using profile id: %s", id)
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
