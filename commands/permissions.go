package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func permissions() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "permissions",
			Aliases:            []string{"perm", "scopes"},
			PostHook:           printJsonAction,
			PersistentPostHook: printCurl,
		},
		getPermissionsCols(),
	)

	listPermissionsCmd(p)
	getPermissionCmd(p)

	return p
}

func getPermissionsCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.permissions.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}
