package commands

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getPermissionCmd(p *commander.Command) *commander.Command {
	gp := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			Aliases:   []string{"check"},
			Example:   "mollie permissions get --permission=customers.write",
			Execute:   getPermissionAction,
			ShortDesc: "Allows the app to check whether an API action is (still) allowed by the authorization.",
			LongDesc: `All API actions through OAuth are by default protected for
privacy and/or money related reasons and therefore require specific permissions.
These permissions can be requested by apps during the OAuth authorization flow.
The Permissions resource allows the app to check whether an API action is (still)
allowed by the authorization.`,
		},
		getPermissionsCols(),
	)

	AddIDFlag(gp, true)

	return gp
}

func getPermissionAction(cmd *cobra.Command, args []string) {
	perm := ParseStringFromFlags(cmd, IDArg)

	res, p, err := app.API.Permissions.Get(
		context.Background(),
		mollie.PermissionGrant(perm),
	)
	if err != nil {
		app.Logger.Fatal(err)
	}

	addStoreValues(Permissions, p, res)

	disp := displayers.MolliePermission{
		Permission: p,
	}

	err = app.Printer.Display(&disp, display.FilterColumns(
		parseFieldsFromFlag(cmd, Permissions),
		getPermissionsCols(),
	))
	if err != nil {
		app.Logger.Fatal(err)
	}
}
