package commands

import (
	"os"

	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ParseStringFromFlags returns the string value of a flag by key.
func ParseStringFromFlags(cmd *cobra.Command, key string) string {
	val, err := cmd.Flags().GetString(key)
	if err != nil {
		logger.Fatal(err)
	}
	return val
}

// ParseIntFromFlags returns the string value of a flag by key.
func ParseIntFromFlags(cmd *cobra.Command, key string) int {
	val, err := cmd.Flags().GetInt(key)
	if err != nil {
		logger.Fatal(err)
	}
	return val
}

// ParsePromptBool returns a boolean to indicate if the values
// should be prompted to the user.
func ParsePromptBool(cmd *cobra.Command) bool {
	val, err := cmd.Flags().GetBool("prompt")
	if err != nil {
		logger.Fatal(err)
	}

	return val
}

// PrintNonemptyFlagValue will log with level info any non empty
// string value.
// The key will be used as name indicator.
// E.g. "using key value: val"
func PrintNonemptyFlagValue(key, val string) {
	if val != "" {
		logger.Infof("using %s value: %s", key, val)
	}
}

// PrintNonEmptyFlags will print all the defined flags for this
// command, both persistent and local flags will be printed.
func PrintNonEmptyFlags(cmd *cobra.Command) {
	cmd.Flags().Visit(printFlagValues)
}

func printFlagValues(f *pflag.Flag) {
	logger.Infof("using %s with value %s", f.Name, f.Value)
}

// PrintJson dumps the given data as json and then it exits
// gracefully from the execution.
func PrintJson(d interface{}) {
	disp := display.Json(d, false)

	printer.Display(disp, commander.NoCols())
	os.Exit(0)
}

// PrintJsonP dumps the given data as pretty json and then it exits
// gracefully from the execution.
func PrintJsonP(d interface{}) {
	disp := display.Json(d, true)

	printer.Display(disp, commander.NoCols())
	os.Exit(0)
}