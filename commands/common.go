package commands

import "github.com/spf13/cobra"

func parseFieldsFromFlag(cmd *cobra.Command) string {
	fields := ParseStringFromFlags(cmd, FieldsArg)

	if fields != "" {
		if verbose {
			PrintNonemptyFlagValue(FieldsArg, fields)
		}
		return fields
	}

	if verbose {
		logger.Info("using all columns (unfiltered)")
	}

	return ""
}
