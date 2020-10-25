package commands

import (
	"time"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/VictorAvelar/mollie-cli/internal/command"
)

const (
	version string = "v0.1.1"
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

	cfgFile          string
	printJSON, debug bool

	// API client
	API *mollie.Client

	// global structured logger
	logger *logrus.Entry
)

func init() {
	MollieCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "specifies a custom config file to be used")
	_ = viper.BindPFlag("mollie.config", MollieCmd.PersistentFlags().Lookup("config"))
	MollieCmd.PersistentFlags().StringVarP(&Token, "token", "t", mollie.APITokenEnv, "the type of token to use for auth")
	_ = viper.BindPFlag("mollie.token", MollieCmd.PersistentFlags().Lookup("token"))
	MollieCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "print verbose logging messages (defaults to false)")
	_ = viper.BindPFlag("mollie.verbose", MollieCmd.PersistentFlags().Lookup("verbose"))
	MollieCmd.PersistentFlags().BoolVar(&printJSON, "print-json", false, "toggle the output type to json")
	_ = viper.BindPFlag("mollie.print-json", MollieCmd.PersistentFlags().Lookup("print-json"))
	MollieCmd.PersistentFlags().StringVarP(&Mode, "mode", "m", string(mollie.TestMode), "indicates the api target from test/live")
	_ = viper.BindPFlag("mode", MollieCmd.PersistentFlags().Lookup("mode"))
	MollieCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enables debug logging information")
	_ = viper.BindPFlag("debug", MollieCmd.PersistentFlags().Lookup("debug"))

	addCommands()
	cobra.OnInitialize(func() {
		initConfig()
		initClient()
	})
}

func initConfig() {
	logger = logrus.WithFields(logrus.Fields{
		"version": version,
		"mode":    Mode,
	})

	if printJSON {
		logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.Logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:        time.RFC822,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			PadLevelText:           true,
		})
	}

	if debug {
		logger.Logger.SetReportCaller(debug)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatal(err)
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
		logger.Fatal(err)
	}

	if Verbose {
		logger.Infof("Using configuration file: %s\n", viper.ConfigFileUsed())
		logger.Infof("Using api token: %s", viper.GetString("mollie.token"))
		logger.Infof("Using api mode: %s", viper.GetString("mollie.mode"))
	}
}

func initClient() {
	var tst bool
	if Mode == string(mollie.LiveMode) {
		tst = !tst
	}

	if Verbose {
		logger.Infof("connecting in %s mode", Mode)
	}

	config := mollie.NewConfig(tst, Token)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		logger.Fatal(err)
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
	MollieCmd.AddCommand(Methods())
}

// ParseStringFromFlags returns the string value of a flag by key.
func ParseStringFromFlags(cmd *cobra.Command, key string) string {
	val, err := cmd.Flags().GetString(key)
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
