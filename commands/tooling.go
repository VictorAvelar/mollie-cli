package commands

import (
	"fmt"

	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// Version creates the version command.
func Version() *command.Command {
	return command.Builder(
		nil,
		command.Config{
			Namespace: "version",
			ShortDesc: "Displays the current command version",
			Execute: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Mollie CLI %s\n", version)
			},
			Aliases: []string{"v", "ver"},
		},
		noCols,
	)
}

// Docs creates the version command.
func Docs() *command.Command {
	return command.Builder(
		nil,
		command.Config{
			Namespace: "docs",
			ShortDesc: "Generates markdown documentation",
			Execute: func(cmd *cobra.Command, args []string) {
				err := doc.GenMarkdownTree(MollieCmd.Command, "./docs")
				if err != nil {
					logger.Fatal(err)
				}
			},
			Hidden: true,
		},
		noCols,
	)
}
