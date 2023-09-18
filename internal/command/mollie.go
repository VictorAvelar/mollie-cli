package command

import "github.com/avocatl/admiral/pkg/commander"

// Cmd returns a command instance.
func Cmd() *commander.Command {
	return commander.Builder(
		nil,
		commander.Config{},
		commander.NoCols(),
	)
}
