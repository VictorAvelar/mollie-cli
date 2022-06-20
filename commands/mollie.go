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

var (
	app *cli

	token, mode, cfgFile       string
	verbose, debug, json, curl bool

	// global structured logger.
	logger *logrus.Entry
)

type cli struct {
	App     *commander.Command
	API     *mollie.Client
	Printer display.Displayer
	Store   map[string]interface{}
	Logger  *logrus.Logger
	Config  *viper.Viper
}

func init() {
	initApp()
}

func initApp() {
	app = &cli{
		App: commander.Builder(
			nil,
			commander.Config{
				Namespace:         "mollie",
				ShortDesc:         "Mollie is a command line interface (CLI) for the Mollie REST API.",
				Version:           version,
				PersistentPreHook: printVerboseFlags,
			},
			commander.NoCols(),
		),
		Logger:  initLogger(),
		Printer: display.DefaultDisplayer(nil),
		Store:   make(map[string]interface{}),
	}

	app.Config = initConfig()
	app.API = initClient()

	addPersistentFlags()
	addCommands()
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:        time.RFC822,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	if debug {
		logger.SetReportCaller(debug)
	}

	return logger
}

func initConfig() *viper.Viper {
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

	if token == "" {
		token = mollie.APITokenEnv
	}

	config := mollie.NewConfig(tst, token)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		app.Logger.Fatal(err)
	}

	return m
}

// Execute runs the command entrypoint.
func Execute() error {
	return app.App.Execute()
}

func addPersistentFlags() {
	commander.AddFlag(app.App, commander.FlagConfig{
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
		app.App.PersistentFlags().Lookup("config"),
	)

	commander.AddFlag(app.App, commander.FlagConfig{
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

	commander.AddFlag(app.App, commander.FlagConfig{
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
		app.App.PersistentFlags().Lookup("mode"),
	)

	commander.AddFlag(app.App, commander.FlagConfig{
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
		app.App.PersistentFlags().Lookup("verbose"),
	)

	commander.AddFlag(app.App, commander.FlagConfig{
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
		app.App.PersistentFlags().Lookup("json"),
	)

	commander.AddFlag(app.App, commander.FlagConfig{
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
		app.App.PersistentFlags().Lookup("curl"),
	)
}

func addCommands() {
	app.App.AddCommand(
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
