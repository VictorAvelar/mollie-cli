package commands

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
)

var (
	profileCols = []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"WEBSITE",
		"EMAIL",
		"PHONE",
		"CATEGORY_CODE",
		"STATUS",
		"REVIEW",
		"CREATED_AT",
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
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve details of a profile, using the profile’s identifier.",
			Execute:   RunGetProfile,
			Example:   "mollie profiles get --id=pfl_token",
		},
		profileCols,
	)

	command.AddFlag(gp, command.FlagConfig{
		Name:     IDArg,
		Usage:    "profile id/token",
		Required: true,
	})

	command.Builder(
		p,
		command.Config{
			Namespace: "current",
			ShortDesc: "Retrieve details of the profile associated to the current API token.",
			LongDesc: `Use this API if you are creating a plugin or SaaS application that allows users to enter a Mollie API key, 
and you want to give a confirmation of the website profile that will be used in your plugin 
or application.`,
			Execute: RunCurrentProfile,
			Example: "mollie profiles current",
		},
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

	if verbose {
		logger.Infof("request target: %s", p.Links.Self.Href)
	}

	mp := displayers.MollieProfile{Profile: p}

	err = command.Display(getProfileCols(cmd), mp.KV())
	if err != nil {
		logger.Error(err)
	}
}

// RunGetProfile will retrieve the required profile details by id.
func RunGetProfile(cmd *cobra.Command, args []string) {
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

	mp := displayers.MollieProfile{Profile: p}

	err = command.Display(getProfileCols(cmd), mp.KV())
	if err != nil {
		logger.Error(err)
	}
}

func getProfileCols(cmd *cobra.Command) []string {
	var cols []string
	{
		cls := ParseStringFromFlags(cmd, FieldsArg)

		if cls != "" {
			cols = strings.Split(cls, ",")
			if verbose {
				PrintNonemptyFlagValue(FieldsArg, cls)
			}
		} else {
			cols = profileCols
			if verbose {
				logger.Info("returning all fields")
			}
		}

	}

	return cols
}
