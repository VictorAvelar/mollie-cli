package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

// MollieProfile wrapper for displaying
type MollieProfile struct {
	*mollie.Profile
}

// KV is a displayable group of key value
func (mp *MollieProfile) KV() []map[string]interface{} {
	out := []map[string]interface{}{}

	x := map[string]interface{}{
		"ID":      mp.Profile.ID,
		"Name":    mp.Profile.Name,
		"Website": mp.Profile.Website,
		"Phone":   mp.Profile.Phone,
		"Status":  mp.Profile.Status,
		"Mode":    mp.Profile.Mode,
		"Since":   mp.CreatedAt.Format("01-02-2006"),
	}

	out = append(out, x)

	return out
}

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

	mp := &MollieProfile{p}

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

	mp := &MollieProfile{p}

	command.Display(profileCols, mp.KV())
}
