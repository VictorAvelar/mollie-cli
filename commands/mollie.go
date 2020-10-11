package commands

import (
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// MollieCmd is the root level mollie-cli command that all other commands attach to
	MollieCmd = &command.Command{
		Command: &cobra.Command{
			Use:   "mollie",
			Short: "Mollie is a command line interface (CLI) for the Mollie REST API.",
		},
	}
	// Token is the main API token
	Token string
	// Mode is the API target (sandbox/live)
	Mode mollie.Mode
	// Verbose toggles verbose output on and off
	Verbose bool

	cfgFile   string
	printJSON bool

	// API client
	API *mollie.Client
)

func init() {
	MollieCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "specifies a custom config file to be used")
	MollieCmd.PersistentFlags().StringVarP(&Token, "token", "t", os.Getenv(mollie.APITokenEnv), "the API token to use (defaults to MOLLIE_API_TOKEN env value)")
	MollieCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "print verbose logging messages (defaults to false)")
	MollieCmd.PersistentFlags().BoolVar(&printJSON, "print-json", false, "toggle the output type to json")

	addCommands()
	cobra.OnInitialize(func() {
		initConfig()
		initClient()
	})
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatal(err)
		}

		viper.SetConfigName(".mollie")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.AddConfigPath(home + "/.config")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	viper.AutomaticEnv()

	if len(viper.GetString("mollie.token")) > 0 {
		Token = viper.GetString("mollie.token")
	}

	if Verbose {
		logrus.Infof("Using configuration file: %s\n", viper.ConfigFileUsed())
		logrus.Infof("Using api token: %s", viper.GetString("mollie.token"))
		logrus.Infof("Using api mode: %s", viper.GetString("mode"))
	}
}

func initClient() {
	config := mollie.NewConfig(true, Token)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		logrus.Fatal(err)
	}

	API = m
}

// Execute runs the command entrypoint
func Execute() error {
	return MollieCmd.Execute()
}

func addCommands() {
	MollieCmd.AddCommand(Profile())
	MollieCmd.AddCommand(Browse())
}
