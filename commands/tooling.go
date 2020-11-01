package commands

import (
	"fmt"

	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// Version creates the version command.
func Version() *command.Command {
	v := command.Builder(
		nil,
		"version",
		"Displays the current command version",
		"",
		func(cmd *cobra.Command, args []string) {
			fmt.Printf("Mollie CLI %s\n", version)
		},
		[]string{},
	)
	v.Aliases = []string{"v", "ver"}

	return v
}

// Docs creates the version command.
func Docs() *command.Command {
	v := command.Builder(
		nil,
		"docs",
		"Generates markdown documentarion",
		"",
		func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(MollieCmd.Command, "./docs")
			if err != nil {
				logger.Fatal(err)
			}
		},
		[]string{},
	)

	v.Hidden = true

	return v
}
