package commands

import (
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	permissionsCols []string = []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"GRANTED",
	}
)

// Permissions builds the permissions command tree.
func Permissions() *command.Command {
	p := command.Builder(
		nil,
		command.Config{
			Namespace: "permissions",
			Aliases:   []string{"perms", "scopes"},
		},
		noCols,
	)

	command.Builder(
		p,
		command.Config{
			Namespace: "list",
			ShortDesc: "List all permissions available with the current app access token.",
			Example:   "mollie permissions list",
			Execute:   RunListPermissions,
		},
		permissionsCols,
	)

	gp := command.Builder(
		p,
		command.Config{
			Namespace: "get",
			Aliases:   []string{"check"},
			Example:   "mollie permissions get --permission=customers.write",
			Execute:   RunGetPermission,
			ShortDesc: "Allows the app to check whether an API action is (still) allowed by the authorization.",
			LongDesc: `All API actions through OAuth are by default protected for
privacy and/or money related reasons and therefore require specific permissions.
These permissions can be requested by apps during the OAuth authorization flow.
The Permissions resource allows the app to check whether an API action is (still)
allowed by the authorization.`,
		},
		paymentsCols,
	)

	command.AddFlag(gp, command.FlagConfig{
		Usage:    "the permission’s ID",
		Required: true,
		Name:     PermissionArg,
	})

	return p
}

// RunListPermissions list all permissions for the current token.
func RunListPermissions(cmd *cobra.Command, args []string) {
	p, err := API.Permissions.List()
	if err != nil {
		logger.Fatal(err)
	}

	disp := displayers.MolliePermissionList{
		PermissionsList: p,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), permissionsCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetPermission allows to check if an API action is (still) allowed by the authorization.
func RunGetPermission(cmd *cobra.Command, args []string) {
	perm := ParseStringFromFlags(cmd, PermissionArg)
	if verbose {
		PrintNonemptyFlagValue(PermissionArg, perm)
	}
	p, err := API.Permissions.Get(perm)
	if err != nil {
		logger.Fatal(err)
	}

	disp := displayers.MolliePermission{
		Permission: p,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), permissionsCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
