package commands

import (
	"time"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	version string = "v0.12.1"

	// Store namespaces
	Payments    = "payments"
	Captures    = "captures"
	Methods     = "methods"
	Profiles    = "profiles"
	Customers   = "customers"
	Permissions = "permissions"
	Chargebacks = "chargebacks"
	Refunds     = "refunds"
	Invoices    = "invoices"
)

type cli struct {
	App     *commander.Command
	API     *mollie.Client
	Printer display.Displayer
	Store   map[string]interface{}
	Logger  *logrus.Logger
	Config  *viper.Viper
}

var (
	app *cli

	token, mode, cfgFile       string
	verbose, debug, json, curl bool

	// global structured logger.
	logger *logrus.Entry
)

func init() {
	app := &cli{}

	app.App = commander.Builder(
		nil,
		commander.Config{
			Namespace:         "mollie",
			ShortDesc:         "Mollie is a command line interface (CLI) for the Mollie REST API.",
			Version:           version,
			PersistentPreHook: printVerboseFlags,
		},
		commander.NoCols(),
	)

	app.Printer = display.DefaultDisplayer(nil)
	app.Config = initConfig()
	app.API = initClient()

	app.addPersistentFlags()
	app.addCommands()
	app.Logger = logrus.WithFields(logrus.Fields{
		"version": version,
		"mode":    mode,
	}).Logger
}

func initConfig() *viper.Viper {
	app.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:        time.RFC822,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	if debug {
		app.Logger.SetReportCaller(debug)
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
		app.Logger.Error(err)
	}

	app.Logger.Printf("%v", viper.AllSettings())

	if verbose {
		app.Logger.Infof("Using configuration file: %s\n", viper.ConfigFileUsed())
		app.Logger.Infof("Using api token: %s", viper.GetString("mollie.token"))
		app.Logger.Infof("Using api mode: %s", viper.GetString("mollie.mode"))
	}

	return viper.GetViper()
}

func initClient() *mollie.Client {
	var tst bool
	if mode == string(mollie.LiveMode) {
		tst = !tst
	}

	if verbose {
		app.Logger.Infof("connecting in %s mode", mode)
	}

	config := mollie.NewConfig(tst, token)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		app.Logger.Fatal(err)
	}

	return m
}

// Execute runs the command entrypoint.
func (a *cli) Execute() error {
	return a.App.Execute()
}

func (a *cli) addPersistentFlags() {
	commander.AddFlag(a.App, commander.FlagConfig{
		Name:       "config",
		Shorthand:  "c",
		Usage:      "specifies a custom config file to be used",
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:      true,
			BindString: &cfgFile,
		},
	})
	_ = viper.BindPFlag(
		"mollie.core.custom_path",
		a.App.PersistentFlags().Lookup("config"),
	)

	commander.AddFlag(a.App, commander.FlagConfig{
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
	_ = viper.BindPFlag(
		"mollie.core.key",
		a.App.PersistentFlags().Lookup("token"),
	)

	commander.AddFlag(a.App, commander.FlagConfig{
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
	_ = viper.BindPFlag(
		"mollie.core.mode",
		a.App.PersistentFlags().Lookup("mode"),
	)

	commander.AddFlag(a.App, commander.FlagConfig{
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
	_ = viper.BindPFlag(
		"mollie.core.verbose",
		a.App.PersistentFlags().Lookup("verbose"),
	)

	commander.AddFlag(a.App, commander.FlagConfig{
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
	_ = viper.BindPFlag(
		"mollie.core.json",
		a.App.PersistentFlags().Lookup("json"),
	)

	commander.AddFlag(a.App, commander.FlagConfig{
		FlagType:   commander.BoolFlag,
		Name:       "curl",
		Usage:      "print the curl representation of a request",
		Default:    false,
		Persistent: true,
		Binding: commander.FlagBindOptions{
			Bound:    true,
			BindBool: &curl,
		},
	})
	_ = viper.BindPFlag("mollie.core.curl",
		a.App.PersistentFlags().Lookup("curl"),
	)
}

func (a *cli) addCommands() {
	a.App.AddCommand(
		docs(),
		browse(),
		methods(),
		permissions(),
		profile(),
		payments(),
		captures(),
		chargebacks(),
		refunds(),
		customers(),
		invoices(),
	)
}
