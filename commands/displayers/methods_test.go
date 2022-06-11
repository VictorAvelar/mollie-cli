package displayers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/stretchr/testify/assert"
)

func TestStringCombinator(t *testing.T) {
	sep := " "
	cases := []struct {
		input  string
		sep    string
		expect string
	}{
		{
			"",
			sep,
			"-",
		},
		{
			"hello world!",
			sep,
			"hello world!",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
			got := stringCombinator(c.sep, strings.Split(c.input, sep)...)
			assert.Equal(t, c.expect, got)
		})
	}
}

func TestMollieMethod_KV(t *testing.T) {
	disp := MollieMethod{
		&mollie.PaymentMethodDetails{
			Resource:      "methods",
			ID:            "ideal",
			Description:   "iDeal payments",
			MinimumAmount: &mollie.Amount{Value: "10.00", Currency: "EUR"},
			MaximumAmount: &mollie.Amount{Value: "100.00", Currency: "EUR"},
			Image: &mollie.Image{
				Size1x: "https://victoravelar.com/logo-example/1.png",
				Size2X: "https://victoravelar.com/logo-example/2x.png",
				Svg:    "https://victoravelar.com/logo-example/logo.svg",
			},
		},
	}

	w := map[string]interface{}{
		"DESCRIPTION": "iDeal payments",
		"ID":          "ideal",
		"LOGO":        "https://victoravelar.com/logo-example/1.png",
		"ISSUERS":     "N/A",
		"MAX_AMOUNT":  "€100.00",
		"MIN_AMOUNT":  "€10.00",
		"RESOURCE":    "methods",
	}

	want := []map[string]interface{}{}
	want = append(want, w)

	assert.Equal(t, want, disp.KV())
}

func TestMollieListMethods(t *testing.T) {
	var meths []*mollie.PaymentMethodDetails
	meths = append(
		meths,
		&mollie.PaymentMethodDetails{
			Resource:      "methods",
			ID:            "ideal",
			Description:   "iDeal payments",
			MinimumAmount: &mollie.Amount{Value: "10.00", Currency: "EUR"},
			MaximumAmount: &mollie.Amount{Value: "100.00", Currency: "EUR"},
			Image: &mollie.Image{
				Size1x: "https://victoravelar.com/logo-example/1.png",
				Size2X: "https://victoravelar.com/logo-example/2x.png",
				Svg:    "https://victoravelar.com/logo-example/logo.svg",
			},
		},
		&mollie.PaymentMethodDetails{
			Resource:      "methods",
			ID:            "paypal",
			Description:   "Paypal",
			MinimumAmount: &mollie.Amount{Value: "10.00", Currency: "EUR"},
			MaximumAmount: nil,
			Image: &mollie.Image{
				Size1x: "https://victoravelar.com/logo-example/1.png",
				Size2X: "https://victoravelar.com/logo-example/2x.png",
				Svg:    "https://victoravelar.com/logo-example/logo.svg",
			},
		})
	disp := MollieListMethods{
		PaymentMethodsList: &mollie.PaymentMethodsList{
			Count: 2,
			Links: mollie.PaginationLinks{
				Documentation: &mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          &mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
			Embedded: struct {
				Methods []*mollie.PaymentMethodDetails
			}{
				Methods: meths,
			},
		},
	}

	var want []map[string]interface{}

	w1 := map[string]interface{}{
		"DESCRIPTION": "iDeal payments",
		"ID":          "ideal",
		"LOGO":        "https://victoravelar.com/logo-example/1.png",
		"ISSUERS":     "N/A",
		"MAX_AMOUNT":  "€100.00",
		"MIN_AMOUNT":  "€10.00",
		"RESOURCE":    "methods",
	}
	w2 := map[string]interface{}{
		"DESCRIPTION": "Paypal",
		"ID":          "paypal",
		"LOGO":        "https://victoravelar.com/logo-example/1.png",
		"ISSUERS":     "N/A",
		"MAX_AMOUNT":  "--- ---",
		"MIN_AMOUNT":  "€10.00",
		"RESOURCE":    "methods",
	}

	want = append(want, w1, w2)

	assert.Equal(t, want, disp.KV())
}
