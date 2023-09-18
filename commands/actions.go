package commands

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"moul.io/http2curl"
)

func printJSONAction(cmd *cobra.Command, args []string) {
	if json {
		ns, ok := app.Store["ns"].(string)
		if !ok {
			log.Fatal("error when matching type")
		}

		data := app.Store[ns]

		printJSONP(data)
	}
}

func printCurl(cmd *cobra.Command, args []string) {
	if curl {
		req, ok := app.Store["request"].(*http.Request)
		if !ok {
			log.Fatal("error when matching type")
		}

		curl, err := http2curl.GetCurlCommand(req)
		if err != nil {
			app.Logger.Error(err)
		}

		app.Logger.Infof(curl.String())
	}
}

func printVerboseFlags(cmd *cobra.Command, args []string) {
	if verbose {
		PrintNonEmptyFlags(cmd)
	}
}
