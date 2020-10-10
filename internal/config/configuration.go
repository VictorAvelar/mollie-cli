package config

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/sirupsen/logrus"
)

// MollieCLIConfiguration describes the config object
// for mollie commands.
type MollieCLIConfiguration struct {
	Token      string      `json:"token,omitempty"`
	Connection mollie.Mode `json:"mode,omitempty"`
	Client     mollie.Client
}

// Initialize will parse the required configuration from
// your config file.
func (mcc *MollieCLIConfiguration) Initialize(logger *logrus.Logger) {
}
