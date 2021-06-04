package commands

import (
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	version string = "v0.12.0"
)

var (
	// MollieCmd is the root level mollie-cli command that all other commands attach to
	MollieCmd = commander.Builder(
		nil,
		commander.Config{
			Namespace: "mollie",
			ShortDesc: "Mollie is a command line interface (CLI) for the Mollie REST API.",
			Version:   version,
		},
		commander.NoCols(),
	)

	token, mode, cfgFile string
	verbose, debug, json bool

	// API client
	API     *mollie.Client
	printer display.Displayer

	// global structured logger
	logger *logrus.Entry
	noCols []string
)

func init() {
	printer = display.DefaultDisplayer(nil)
	addPersistentFlags()
	addCommands()
	cobra.OnInitialize(func() {
		initConfig()
		initClient()
	})
}

func initConfig() {
	logger = logrus.WithFields(logrus.Fields{
		"version": version,
		"mode":    mode,
	})

	logger.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:        time.RFC822,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

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

	if verbose {
		logger.Infof("Using configuration file: %s\n", viper.ConfigFileUsed())
		logger.Infof("Using api token: %s", viper.GetString("mollie.token"))
		logger.Infof("Using api mode: %s", viper.GetString("mollie.mode"))
	}
}

func initClient() {
	var tst bool
	if mode == string(mollie.LiveMode) {
		tst = !tst
	}

	if verbose {
		logger.Infof("connecting in %s mode", mode)
	}

	config := mollie.NewConfig(tst, token)
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

func addPersistentFlags() {
	commander.AddFlag(MollieCmd, commander.FlagConfig{
		Name:       "config",
		Shorthand:  "c",
		Usage:      "specifies a custom config file to be used",
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:      true,
			BindString: &cfgFile,
		},
	})
	_ = viper.BindPFlag("mollie.config", MollieCmd.PersistentFlags().Lookup("config"))
	commander.AddFlag(MollieCmd, commander.FlagConfig{
		Name:       "token",
		Shorthand:  "t",
		Usage:      "the type of token to use for auth",
		Default:    mollie.APITokenEnv,
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:      true,
			BindString: &token,
		},
	})
	_ = viper.BindPFlag("mollie.token", MollieCmd.PersistentFlags().Lookup("token"))
	commander.AddFlag(MollieCmd, commander.FlagConfig{
		Name:       "mode",
		Shorthand:  "m",
		Usage:      "indicates the api target from test/live",
		Default:    string(mollie.TestMode),
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:      true,
			BindString: &mode,
		},
	})
	_ = viper.BindPFlag("mode", MollieCmd.PersistentFlags().Lookup("mode"))
	commander.AddFlag(MollieCmd, commander.FlagConfig{
		FlagType:   commander.BoolFlag,
		Name:       "verbose",
		Shorthand:  "v",
		Usage:      "print verbose logging messages (defaults to false)",
		Default:    false,
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:    true,
			BindBool: &verbose,
		},
	})
	_ = viper.BindPFlag("mollie.verbose", MollieCmd.PersistentFlags().Lookup("verbose"))
	commander.AddFlag(MollieCmd, commander.FlagConfig{
		FlagType:   commander.BoolFlag,
		Name:       "debug",
		Shorthand:  "d",
		Usage:      "enables debug logging information",
		Default:    false,
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:    true,
			BindBool: &debug,
		},
	})
	_ = viper.BindPFlag("debug", MollieCmd.PersistentFlags().Lookup("debug"))

	commander.AddFlag(MollieCmd, commander.FlagConfig{
		FlagType:   commander.BoolFlag,
		Name:       "json",
		Usage:      "dumpts the json response instead of the column based output",
		Default:    false,
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:    true,
			BindBool: &json,
		},
	})
	_ = viper.BindPFlag("json", MollieCmd.PersistentFlags().Lookup("json"))
}

func addCommands() {
	MollieCmd.AddCommand(
		docs(),
		browse(),
		methods(),
		permissions(),
		profile(),
	)
	// MollieCmd.AddCommand(Profile())
	// MollieCmd.AddCommand(Methods())
	// MollieCmd.AddCommand(Payments())
	// MollieCmd.AddCommand(Chargebacks())
	// MollieCmd.AddCommand(Refunds())
	// MollieCmd.AddCommand(Customers())
	// MollieCmd.AddCommand(Captures())
	// MollieCmd.AddCommand(Permissions())
	// MollieCmd.AddCommand(Invoices())

	// // Tooling
	// MollieCmd.AddCommand(Version())
	// MollieCmd.AddCommand(Docs())
}
