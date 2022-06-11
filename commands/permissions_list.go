package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func listPermissionsCmd(p *commander.Command) *commander.Command {
	return commander.Builder(
		p,
		commander.Config{
			Namespace: "list",
			ShortDesc: "List all permissions available with the current app access token.",
			Example:   "mollie permissions list",
			Execute:   listPermissionsAction,
		},
		getPermissionsCols(),
	)
}

// RunListPermissions list all permissions for the current token.
func listPermissionsAction(cmd *cobra.Command, args []string) {
	_, p, err := API.Permissions.List(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	disp := displayers.MolliePermissionList{
		PermissionsList: p,
	}

	err = printer.Display(&disp, display.FilterColumns(
		parseFieldsFromFlag(cmd),
		getPermissionsCols(),
	))

	if err != nil {
		logger.Fatal(err)
	}
}
