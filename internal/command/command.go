package command

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Command wraps a base cobra command to add some
// custom functionality.
type Command struct {
	*cobra.Command
	Cols     []string
	children []*Command
}

// AddCommand adds child commands and also to cobra.
func (c *Command) AddCommand(commands ...*Command) {
	c.children = append(c.children, commands...)
	for _, cmd := range commands {
		c.Command.AddCommand(cmd.Command)
	}
}

// GetSubCommands returns the command sub commands.
func (c *Command) GetSubCommands() []*Command {
	return c.children
}

// Builder constructs a new command.
func Builder(parent *Command, config Config, cols []string) *Command {
	cc := &cobra.Command{
		Use:       config.Namespace,
		Short:     config.ShortDesc,
		Long:      strings.TrimSpace(config.LongDesc),
		Run:       config.Execute,
		PreRun:    config.PreHook,
		PostRun:   config.PostHook,
		Hidden:    config.Hidden,
		ValidArgs: config.ValidArgs,
		Example:   config.Example,
		Aliases:   config.Aliases,
	}

	c := &Command{Command: cc, Cols: cols}

	if parent != nil {
		parent.AddCommand(c)
	}

	return c
}

// Display pretty prints a tab version of the data.
func Display(c []string, vals []map[string]interface{}) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 4, ' ', 0)

	fmt.Fprintln(w, strings.Join(c, "\t"))

	for _, r := range vals {
		values := []interface{}{}
		formats := []string{}

		for _, col := range c {
			v := r[col]
			values = append(values, v)

			switch v.(type) {
			case string:
				formats = append(formats, "%s")
			case int:
				formats = append(formats, "%d")
			case float64:
				formats = append(formats, "%f")
			case bool:
				formats = append(formats, "%v")
			default:
				formats = append(formats, "%v")
			}
		}
		format := strings.Join(formats, "\t")
		fmt.Fprintf(w, format+"\n", values...)
	}
	return w.Flush()
}

// AddFlag attaches a flag of the given type with the
// specified configuration.
func AddFlag(cmd *Command, config FlagConfig) {
	var flagger *pflag.FlagSet
	{
		if config.Persistent {
			flagger = cmd.PersistentFlags()
		} else {
			flagger = cmd.Flags()
		}
	}
	switch config.FlagType {
	case IntFlag:
		val := config.Default.(int)
		flagger.IntP(config.Name, config.Shorthand, val, config.Usage)
	case Int64Flag:
		val := config.Default.(int64)
		flagger.Int64P(config.Name, config.Shorthand, val, config.Usage)
	case Float64Flag:
		val := config.Default.(float64)
		flagger.Float64P(config.Name, config.Shorthand, val, config.Usage)
	case BoolFlag:
		val := config.Default.(bool)
		flagger.BoolP(config.Name, config.Shorthand, val, config.Usage)
	default:
		config.Default = ""
		val := config.Default.(string)
		flagger.StringP(config.Name, config.Shorthand, val, config.Usage)
	}

	if config.Required {
		err := cmd.MarkFlagRequired(config.Name)
		if err != nil {
			logrus.Error(err)
		}
	}
}
