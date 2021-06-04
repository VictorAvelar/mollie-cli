package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func docs() *commander.Command {
	return commander.Builder(
		nil,
		commander.Config{
			Namespace: "docs",
			ShortDesc: "Generates markdown documentation",
			Execute: func(cmd *cobra.Command, args []string) {
				err := doc.GenMarkdownTree(MollieCmd.Command, "./docs")
				if err != nil {
					logger.Fatal(err)
				}
			},
			Hidden:  true,
			Example: "mollie docs",
		},
		commander.NoCols(),
	)
}
