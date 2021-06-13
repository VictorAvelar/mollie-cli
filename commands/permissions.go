package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func permissions() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace: "permissions",
			Aliases:   []string{"perm", "scopes"},
		},
		getPermissionsCols(),
	)

	listPermissionsCmd(p)
	getPermissionCmd(p)

	return p
}

func getPermissionsCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"GRANTED",
	}
}
