package commands

import (
	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
)

var (
	profileCols = []string{
		"ID",
		"Name",
		"Website",
		"Phone",
		"Status",
		"Mode",
		"Since",
	}
)

// Profile creates the profile commands tree.
func Profile() *command.Command {
	p := &command.Command{
		Command: &cobra.Command{
			Use:   "profiles",
			Short: "In order to process payments, you need to create a website profile",
		},
	}

	gp := command.Builder(
		p,
		"get",
		"Retrieve details of a profile, using the profile’s identifier.",
		"",
		RunGetProfile,
		profileCols,
	)

	command.AddStringFlag(gp, IDArg, "", "", "profile ID to be retrieved", true)

	command.Builder(
		p,
		"current",
		"Retrieve details of the profile associated to the current API token.",
		`Use this API if you are creating a plugin or SaaS application that allows users to enter a Mollie API key, 
and you want to give a confirmation of the website profile that will be used in your plugin 
or application.`,
		RunCurrentProfile,
		profileCols,
	)

	return p
}

// RunCurrentProfile executes the get current profile action.
func RunCurrentProfile(cmd *cobra.Command, args []string) {
	p, err := API.Profiles.Current()
	if err != nil {
		logger.Fatal(err)
	}

	if Verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
	}

	mp := displayers.MollieProfile{Profile: p}

	err = command.Display(profileCols, mp.KV())
	if err != nil {
		logger.Error(err)
	}
}

// RunGetProfile will retrieve the required profile details by id.
func RunGetProfile(cmd *cobra.Command, args []string) {
	id := ParseStringFromFlags(cmd, IDArg)

	if Verbose {
		logger.Infof("fetching profile with id %s", id)
	}

	p, err := API.Profiles.Get(id)
	if err != nil {
		logger.Fatal(err)
	}
	if Verbose {
		logger.Infof("using profile id: %s", id)
		logger.Infof("request target: %s", p.Links.Self.Href)
	}

	mp := displayers.MollieProfile{Profile: p}

	err = command.Display(profileCols, mp.KV())
	if err != nil {
		logger.Error(err)
	}
}
