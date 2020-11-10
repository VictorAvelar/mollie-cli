package displayers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
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
		&mollie.PaymentMethodInfo{
			ID:            "ideal",
			Description:   "iDeal payments",
			MinimumAmount: &mollie.Amount{Value: "10.00", Currency: "EUR"},
			MaximumAmount: &mollie.Amount{Value: "100.00", Currency: "EUR"},
		},
	}

	w := map[string]interface{}{
		"ID":             "ideal",
		"Name":           "iDeal payments",
		"Minimum Amount": "10.00/EUR",
		"Maximum Amount": "100.00/EUR",
	}

	want := []map[string]interface{}{}
	want = append(want, w)

	assert.Equal(t, want, disp.KV())
}

func TestMollieListMethods(t *testing.T) {
	var meths []*mollie.PaymentMethodInfo
	meths = append(
		meths,
		&mollie.PaymentMethodInfo{
			ID:            "ideal",
			Description:   "iDeal payments",
			MinimumAmount: &mollie.Amount{Value: "10.00", Currency: "EUR"},
			MaximumAmount: &mollie.Amount{Value: "100.00", Currency: "EUR"},
		},
		&mollie.PaymentMethodInfo{
			ID:            "paypal",
			Description:   "PayPal",
			MinimumAmount: &mollie.Amount{Value: "0.01", Currency: "EUR"},
			MaximumAmount: &mollie.Amount{Value: "", Currency: ""},
		})
	disp := MollieListMethods{
		ListMethods: &mollie.ListMethods{
			Count: 2,
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
			Embedded: struct{ Methods []*mollie.PaymentMethodInfo }{
				Methods: meths,
			},
		},
	}

	var want []map[string]interface{}

	w1 := map[string]interface{}{
		"ID":             "ideal",
		"Name":           "iDeal payments",
		"Minimum Amount": "10.00/EUR",
		"Maximum Amount": "100.00/EUR",
	}
	w2 := map[string]interface{}{
		"ID":             "paypal",
		"Name":           "PayPal",
		"Minimum Amount": "0.01/EUR",
		"Maximum Amount": "-/-",
	}

	want = append(want, w1, w2)

	assert.Equal(t, want, disp.KV())
}
