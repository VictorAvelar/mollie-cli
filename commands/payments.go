package commands

import (
	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/prompter"
)

func payments() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace:          "payments",
			ShortDesc:          "All operations to handle payments",
			Aliases:            []string{"pay", "p"},
			PostHook:           printJsonAction,
			PersistentPostHook: printCurl,
		},
		getPaymentCols(),
	)

	listPaymentsCmd(p)
	getPaymentCmd(p)
	cancelPaymentCmd(p)
	createPaymentCmd(p)
	updatePaymentCmd(p)

	return p
}

func getPaymentCols() []string {
	cols := app.Config.GetStringSlice("mollie.fields.payments.all")

	if verbose {
		app.Logger.Infof("parsed fields %v", cols)
	}

	return cols
}

func attachPaymentMethodSpecificValues(p *mollie.Payment) {
	switch p.Method {
	case mollie.BankTransfer:
		p.BillingEmail = promptStringClean("billing email address", "")
		p.DueDate = promptShortDate()
	case mollie.CreditCard:
		p.BillingAddress = promptAddress()
		p.CardToken = promptStringClean("card token", "")
		p.ShippingAddress = promptPaymentDetailsAddress()
	case mollie.GiftCard:
		p.Issuer = promptGiftCardIssuer()
		p.VoucherNumber = promptStringClean("voucher number", "")
		p.VoucherPin = promptStringClean("voucher pin", "")
	case mollie.IDeal:
		p.Issuer = promptStringClean("ideal issuer", "")
	case mollie.KBC:
		p.Issuer = promptKbcIssuer()
	case mollie.KlarnaPayLater, mollie.KlarnaLiceit:
		app.Logger.Fatal("for the selected payment method you need to use the orders api")
	case mollie.PayPal:
		p.ShippingAddress = promptPaymentDetailsAddress()
		p.SessionID = promptStringClean("session id", "")
		if err := prompter.Confirm("is a digital good", nil); err == nil {
			p.DigitalGoods = true
		}
	case mollie.PaySafeCard:
		p.CustomerReference = promptStringClean("customer reference", "")
	case mollie.PRZelewy24:
		p.BillingEmail = promptStringClean("billing email address", "")
	case mollie.DirectDebit:
		p.ConsumerName = promptStringClean("consumer name", "")
		p.ConsumerAccount = promptStringClean("consumer account", "")
	case mollie.Bancontact, mollie.Belfius, mollie.EPS, mollie.GiroPay, mollie.MyBank, mollie.Sofort:
		if verbose {
			app.Logger.Info("there are no payment method specific fields for your selection")
		}
	}
}

func attachAccessTokenParams(p *mollie.Payment) *mollie.Payment {
	if app.API.HasAccessToken() {
		p.ProfileID = promptStringClean("profile id", "")
	}

	return p
}
