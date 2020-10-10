package runners

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Runner describe cobra executable functions.
type Runner func(cmd *cobra.Command, args []string)

// NopRunner is an empty no operation runner.
func NopRunner(cmd *cobra.Command, args []string) {}

// ArgPrinterRunner will print all the received arguments.
func ArgPrinterRunner(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
