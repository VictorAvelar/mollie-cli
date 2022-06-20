package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func parseFieldsFromFlag(cmd *cobra.Command, r string) string {
	fields := ParseStringFromFlags(cmd, FieldsArg)

	if fields != "" {
		if verbose {
			PrintNonemptyFlagValue(FieldsArg, fields)
		}
		return fields
	}

	cfgFields := strings.Join(
		app.Config.GetStringSlice(
			fmt.Sprintf("mollie.fields.%s.printable", r)),
		",",
	)

	if cfgFields != "" {
		if verbose {
			app.Logger.Info("using fields from config file")
		}

		return cfgFields
	}

	if verbose {
		app.Logger.Info("using all columns (unfiltered)")
	}

	return ""

}
