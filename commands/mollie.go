package commands

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/VictorAvelar/mollie-cli/internal/command"
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
	Mode string
	// Verbose toggles verbose output on and off
	Verbose bool

	cfgFile   string
	printJSON bool

	// API client
	API *mollie.Client
)

func init() {
	MollieCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "specifies a custom config file to be used")
	viper.BindPFlag("mollie.config", MollieCmd.PersistentFlags().Lookup("config"))
	MollieCmd.PersistentFlags().StringVarP(&Token, "token", "t", mollie.APITokenEnv, "the type of token to use for auth (defaults to MOLLIE_API_TOKEN)")
	viper.BindPFlag("mollie.token", MollieCmd.PersistentFlags().Lookup("token"))
	MollieCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "print verbose logging messages (defaults to false)")
	viper.BindPFlag("mollie.verbose", MollieCmd.PersistentFlags().Lookup("verbose"))
	MollieCmd.PersistentFlags().BoolVar(&printJSON, "print-json", false, "toggle the output type to json")
	viper.BindPFlag("mollie.print-json", MollieCmd.PersistentFlags().Lookup("print-json"))
	MollieCmd.PersistentFlags().StringVarP(&Mode, "mode", "m", string(mollie.TestMode), "indicates the api target from test/live")
	viper.BindPFlag("mode", MollieCmd.PersistentFlags().Lookup("mode"))

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

		viper.SetEnvPrefix("MOLLIE")
		viper.AutomaticEnv()
		viper.SetConfigName(".mollie")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.AddConfigPath(home + "/.config")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if Verbose {
		logrus.Infof("Using configuration file: %s\n", viper.ConfigFileUsed())
		logrus.Infof("Using api token: %s", viper.GetString("mollie.token"))
		logrus.Infof("Using api mode: %s", viper.GetString("mollie.mode"))
	}
}

func initClient() {
	var tst bool
	if Mode == string(mollie.LiveMode) {
		tst = !tst
	}

	if Verbose {
		logrus.Infof("connecting in %s mode", Mode)
	}

	config := mollie.NewConfig(tst, Token)
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
