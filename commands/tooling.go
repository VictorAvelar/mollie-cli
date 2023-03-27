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
				if len(args) == 0 {
					err := doc.GenMarkdownTree(app.App.Command, "./docs")
					if err != nil {
						app.Logger.Fatal(err)
					}
				}

				if args[0] == "man" {
					err := doc.GenManTree(app.App.Command, nil, "./manpages")
					if err != nil {
						app.Logger.Fatal(err)
					}
				}
			},
			Hidden:  true,
			Example: "mollie docs",
		},
		commander.NoCols(),
	)
}
