package command

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/VictorAvelar/mollie-cli/internal/runners"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
func Builder(parent *Command, cliText, shortDesc, desc string, cr runners.Runner, cols []string) *Command {
	cc := &cobra.Command{
		Use:   cliText,
		Short: shortDesc,
		Long:  strings.TrimSpace(desc),
		Run:   cr,
	}

	c := &Command{Command: cc, Cols: cols}

	if parent != nil {
		parent.AddCommand(c)
	}

	if len(cols) > 0 {
		cc.Flags().StringP("props", "p", strings.Join(cols, ","), "List of properties to display")
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

// AddStringFlag will decorate the command with the string flag requirements.
func AddStringFlag(cmd *Command, name, short, fb, usage string, req bool) {
	cmd.Flags().StringP(name, short, fb, usage)
	if req {
		err := cmd.MarkFlagRequired(name)
		if err != nil {
			logrus.Error(err)
		}
	}
}

// AddBoolFlag will decorate the command with the bool flag requirements.
func AddBoolFlag(cmd *Command, name, short, usage string, fb, req bool) {
	cmd.Flags().BoolP(name, short, fb, usage)
	if req {
		err := cmd.MarkFlagRequired(name)
		if err != nil {
			logrus.Error(err)
		}
	}
}

// AddIntFlag will decorate the command with the int flag requirements.
func AddIntFlag(cmd *Command, name, short, usage string, fb int, req bool) {
	cmd.Flags().IntP(name, short, fb, usage)
	if req {
		err := cmd.MarkFlagRequired(name)
		if err != nil {
			logrus.Error(err)
		}
	}
}

// AddInt64Flag will decorate the command with the int64 flag requirements.
func AddInt64Flag(cmd *Command, name, short, usage string, fb int64, req bool) {
	cmd.Flags().Int64P(name, short, fb, usage)
	if req {
		err := cmd.MarkFlagRequired(name)
		if err != nil {
			logrus.Error(err)
		}
	}
}
