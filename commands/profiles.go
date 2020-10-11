package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
)

const (
	profileIDArgName string = "id"
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

	command.AddStringFlag(gp, profileIDArgName, "", "", "profile ID to be retrieved", true)

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
		logrus.Fatal(err)
	}

	mp := displayers.MollieProfile{Profile: p}

	command.Display(profileCols, mp.KV())
}

// RunGetProfile will retrieve the required profile details by id.
func RunGetProfile(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetString(profileIDArgName)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("using profile id: %s", id)

	p, err := API.Profiles.Get(id)
	if err != nil {
		logrus.Fatal(err)
	}

	mp := displayers.MollieProfile{Profile: p}

	command.Display(profileCols, mp.KV())
}
