package command

import "github.com/avocatl/admiral/pkg/commander"

func Cmd() *commander.Command {
	return commander.Builder(
		nil,
		commander.Config{},
		commander.NoCols(),
	)
}
