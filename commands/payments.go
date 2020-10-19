package commands

import (
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

// Payments builds the payments command tree.
func Payments() *command.Command {
	p := &command.Command{
		Command: &cobra.Command{
			Use:     "payments",
			Short:   "All operations to handle payments",
			Aliases: []string{"pay", "p"},
		},
	}

	return p
}
