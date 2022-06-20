package commands

import (
	"net/http"

	"github.com/spf13/cobra"
	"moul.io/http2curl"
)

func printJsonAction(cmd *cobra.Command, args []string) {
	if json {
		ns := app.Store["ns"].(string)
		data := app.Store[ns]

		printJSONP(data)
	}
}

func printCurl(cmd *cobra.Command, args []string) {
	if curl {
		req := app.Store["request"].(*http.Request)

		curl, err := http2curl.GetCurlCommand(req)
		if err != nil {
			app.Logger.Error(err)
		}

		app.Logger.Infof(`
Your cURL:
%s
`, curl)
	}
}

func printVerboseFlags(cmd *cobra.Command, args []string) {
	if verbose {
		PrintNonEmptyFlags(cmd)
	}
}
