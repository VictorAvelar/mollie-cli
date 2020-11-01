package command

import "github.com/VictorAvelar/mollie-cli/internal/runners"

// Config contains a command configuration
type Config struct {
	Aliases   []string
	Example   string
	Execute   runners.Runner
	Hidden    bool
	LongDesc  string
	Namespace string
	PostHook  runners.Runner
	PreHook   runners.Runner
	ShortDesc string
	ValidArgs []string
}

// Supported flags
const (
	StringFlag = iota
	IntFlag
	Int64Flag
	Float64Flag
	BoolFlag
)

// FlagConfig defines the configuration of a flag.
type FlagConfig struct {
	FlagType   int
	Name       string
	Shorthand  string
	Usage      string
	Default    interface{}
	Required   bool
	Persistent bool
}
