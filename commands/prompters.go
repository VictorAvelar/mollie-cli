package commands

import (
	"context"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/avocatl/admiral/pkg/prompter"
)

func promptPaymentMethod() mollie.PaymentMethod {
	_, methods, err := app.API.PaymentMethods.List(context.Background(), nil)
	if err != nil {
		app.Logger.Fatal(err)
	}

	var ms []string
	{
		for _, m := range methods.Embedded.Methods {
			ms = append(ms, m.ID)
		}
	}

	v, err := prompter.Select("Select a payment method:", ms)
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(string)

	return mollie.PaymentMethod(val)
}

func promptAddress() *mollie.Address {
	v, err := prompter.Struct(&mollie.Address{})
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(*mollie.Address)

	return val
}

func promptAmount() *mollie.Amount {
	v, err := prompter.Struct(&mollie.Amount{})
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(*mollie.Amount)

	return val
}

func promptSequenceType() mollie.SequenceType {
	v, err := prompter.Select("sequence type:", []string{
		string(mollie.OneOffSequence),
		string(mollie.FirstSequence),
		string(mollie.RecurringSequence),
	})

	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(string)

	return mollie.SequenceType(val)
}

func promptShortDate() *mollie.ShortDate {
	v := promptStringClean("due date (YYYY-MM-DD)", time.Now().AddDate(0, 0, 1).Format("2006-01-02"))

	t, err := time.Parse("2006-01-02", v)
	if err != nil {
		app.Logger.Fatal(err)
	}

	return &mollie.ShortDate{
		Time: t,
	}
}

func promptPaymentDetailsAddress() *mollie.PaymentDetailsAddress {
	v, err := prompter.Struct(&mollie.PaymentDetailsAddress{})
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(*mollie.PaymentDetailsAddress)

	return val
}

func promptGiftCardIssuer() string {
	v, err := prompter.Select("gift card issuer", []string{
		string(mollie.BloemenCadeuKaart),
		string(mollie.Boekenbon),
		string(mollie.DecaudeuKaart),
		string(mollie.DelokaleDecauKaart),
		string(mollie.Dinercadeau),
		string(mollie.Fashioncheque),
		string(mollie.Festivalcadeau),
		string(mollie.Good4fun),
		string(mollie.KlusCadeu),
		string(mollie.Kunstencultuurcadeaukaart),
		string(mollie.Nationalebioscoopbon),
		string(mollie.Nationaleentertainmentcard),
		string(mollie.Nationalegolfbon),
		string(mollie.Ohmygood),
		string(mollie.Podiumcadeaukaart),
		string(mollie.Reiscadeau),
		string(mollie.Restaurantcadeau),
		string(mollie.Sportenfitcadeau),
		string(mollie.Sustainablefashion),
		string(mollie.Travelcheq),
		string(mollie.Vvvgiftcard),
		string(mollie.Vvvdinercheque),
		string(mollie.Vvvlekkerweg),
		string(mollie.Webshopgiftcard),
		string(mollie.Yourgift),
	})
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(string)

	return val
}

func promptLocale() mollie.Locale {
	v, err := prompter.Select("Locale for your payment?", getMollieLocales())
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(string)

	return mollie.Locale(val)
}

func promptKbcIssuer() string {
	v, err := prompter.Select("Locale for your payment?", []string{
		"kbc",
		"cbc",
	})
	if err != nil {
		app.Logger.Fatal(err)
	}

	val := v.(string)

	return val
}

func getMollieLocales() []string {
	return []string{
		string(mollie.English),
		string(mollie.Dutch),
		string(mollie.DutchBelgium),
		string(mollie.French),
		string(mollie.FrenchBelgium),
		string(mollie.German),
		string(mollie.GermanAustria),
		string(mollie.GermanSwiss),
		string(mollie.Spanish),
		string(mollie.Catalan),
		string(mollie.Portuguese),
		string(mollie.Italian),
		string(mollie.Norwegian),
		string(mollie.Swedish),
		string(mollie.Finish),
		string(mollie.Danish),
		string(mollie.Icelandic),
		string(mollie.Hungarian),
		string(mollie.Polish),
		string(mollie.Latvian),
		string(mollie.Lithuanian),
	}
}
