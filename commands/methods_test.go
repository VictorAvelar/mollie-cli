package commands

import (
	"os"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func TestRunGetMethods(t *testing.T) {
	os.Setenv("MOLLIE_API_TOKEN", "test_9yr4WmkhQ38DCm2majrPPVx92PtpvU")
	initClient()
	cmd := &cobra.Command{}

	cmd.Flags().String("locale", "", "")
	cmd.Flags().String("id", "", "")
	cmd.MarkFlagRequired("id")
	cmd.Flags().String("currency", "", "")

	err := cmd.Flags().Set("locale", "en_US")
	if err != nil {
		t.Error(err)
	}
	err = cmd.Flags().Set("currency", "USD")
	if err != nil {
		t.Error(err)
	}
	err = cmd.Flags().Set("id", string(mollie.IDeal))
	if err != nil {
		t.Error(err)
	}

	val, _ := cmd.Flags().GetString("locale")
	logrus.Infof("%+v", val)

	RunGetPaymentMethods(cmd, []string{})
}
