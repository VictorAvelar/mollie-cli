package commands

import (
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/VictorAvelar/mollie-cli/internal/command"
	"github.com/spf13/cobra"
)

var (
	captureCols = []string{
		"RESOURCE",
		"ID",
		"MODE",
		"AMOUNT",
		"SETTLEMENT_AMOUNT",
		"PAYMENT_ID",
		"SHIPMENT_ID",
		"SETTLEMENT_ID",
		"CREATED_AT",
	}
)

// Captures creates the captures commands tree.
func Captures() *command.Command {
	c := command.Builder(
		nil,
		command.Config{
			Namespace: "captures",
			ShortDesc: "Operations with Captures API.",
		},
		noCols,
	)

	lc := command.Builder(
		c,
		command.Config{
			Namespace: "list",
			ShortDesc: "Retrieve all captures for a certain payment.",
			LongDesc: `Retrieve all captures for a certain payment.
Captures are used for payments that have the authorize-then-capture flow. 
The only payment methods at the moment that have this flow are Klarna Pay 
later and Klarna Slice it.`,
			Execute: RunListCaptures,
			Example: "mollie captures list --payment tr_example",
		},
		captureCols,
	)

	command.AddFlag(lc, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "the payment id/token",
		Required: true,
	})

	gc := command.Builder(
		c,
		command.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single capture by its ID.",
			LongDesc: `Retrieve a single capture by its ID. Note the original payment’s ID is needed as well.
Captures are used for payments that have the authorize-then-capture flow. 
The only payment methods at the moment that have this flow are Klarna Pay 
later and Klarna Slice it.`,
			Execute: RunGetCapture,
			Example: "mollie captures get --payment tr_example --id ct_example",
		},
		captureCols,
	)

	command.AddFlag(gc, command.FlagConfig{
		Name:     PaymentArg,
		Usage:    "the payment id/token",
		Required: true,
	})

	command.AddFlag(gc, command.FlagConfig{
		Name:     IDArg,
		Usage:    "the capture id/token",
		Required: true,
	})

	return c
}

// RunListCaptures will return all captures for a specified payment.
func RunListCaptures(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)

	if verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
	}

	captures, err := API.Captures.List(payment)
	if err != nil {
		logger.Fatal(err)
	}

	if verbose {
		logger.Infof("request target: %s", captures.Links.Self.Href)
		logger.Infof("request docs: %s", captures.Links.Documentation.Href)
	}

	disp := displayers.MollieCapturesList{
		CapturesList: captures,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), captureCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}

// RunGetCapture retrieves a capture associated to a payment
// by it id/token.
func RunGetCapture(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(IDArg, id)
	}

	capture, err := API.Captures.Get(payment, id)
	if err != nil {
		logger.Fatal(err)
	}

	disp := displayers.MollieCapture{
		Capture: capture,
	}

	err = command.Display(
		command.FilterColumns(parseFieldsFromFlag(cmd), captureCols),
		disp.KV(),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
