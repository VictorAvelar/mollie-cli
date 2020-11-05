package displayers

import (
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMolliePayment_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "04-11-2020")
	if err != nil {
		t.Error(err)
	}
	disp := MolliePayment{
		Payment: &mollie.Payment{
			ID:            "tr_test",
			Mode:          mollie.TestMode,
			CreatedAt:     n,
			ExpiresAt:     n.AddDate(0, 0, 2),
			IsCancellable: false,
			Amount:        mollie.Amount{Currency: "EUR", Value: "1.00"},
			Method:        mollie.PayPal,
			Description:   "testing KV",
		},
	}

	w := map[string]interface{}{
		"ID":          "tr_test",
		"Mode":        mollie.TestMode,
		"Created":     n.Format("02-01-2006"),
		"Expires":     n.AddDate(0, 0, 2).Format("02-01-2006"),
		"Cancelable":  false,
		"Amount":      "1.00 EUR",
		"Method":      "paypal",
		"Description": "testing KV",
	}

	want := []map[string]interface{}{}
	want = append(want, w)

	assert.Equal(t, want, disp.KV())
}

func TestMollieListPayments_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "04-11-2020")
	if err != nil {
		t.Error(err)
	}
	var ps []mollie.Payment
	{
		ps = append(
			ps,
			mollie.Payment{
				ID:            "tr_test",
				Mode:          mollie.TestMode,
				CreatedAt:     n,
				ExpiresAt:     n.AddDate(0, 0, 2),
				IsCancellable: false,
				Amount:        mollie.Amount{Currency: "EUR", Value: "1.00"},
				Method:        mollie.PayPal,
				Description:   "testing KV",
			},
			mollie.Payment{
				ID:            "tr_test_2",
				Mode:          mollie.TestMode,
				CreatedAt:     n,
				ExpiresAt:     n.AddDate(0, 0, 2),
				IsCancellable: false,
				Amount:        mollie.Amount{Currency: "USD", Value: "2.00"},
				Method:        mollie.BankTransfer,
				Description:   "testing KV list payments",
			},
		)
	}

	disp := MollieListPayments{
		PaymentList: &mollie.PaymentList{
			Count: 2,
			Embedded: struct{ Payments []mollie.Payment }{
				Payments: ps,
			},
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	w := map[string]interface{}{
		"ID":          "tr_test",
		"Mode":        mollie.TestMode,
		"Created":     n.Format("02-01-2006"),
		"Expires":     n.AddDate(0, 0, 2).Format("02-01-2006"),
		"Cancelable":  false,
		"Amount":      "1.00 EUR",
		"Method":      "paypal",
		"Description": "testing KV",
	}

	w1 := map[string]interface{}{
		"ID":          "tr_test_2",
		"Mode":        mollie.TestMode,
		"Created":     n.Format("02-01-2006"),
		"Expires":     n.AddDate(0, 0, 2).Format("02-01-2006"),
		"Cancelable":  false,
		"Amount":      "2.00 USD",
		"Method":      "banktransfer",
		"Description": "testing KV list payments",
	}

	want := []map[string]interface{}{}
	want = append(want, w, w1)

	assert.Equal(t, want, disp.KV())
}

func TestGetSafeExpiration(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	cases := []struct {
		name   string
		expect string
		input  mollie.Payment
	}{
		{
			name:   "non zero date",
			expect: "01-11-2020",
			input: mollie.Payment{
				ExpiresAt: n,
			},
		},
		{
			name:   "zero date",
			expect: "----------",
			input:  mollie.Payment{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := getSafeExpiration(c.input)
			assert.Equal(t, c.expect, got)
		})
	}
}

func TestGetSafePaymentMethod(t *testing.T) {
	cases := []struct {
		name   string
		expect string
		input  mollie.Payment
	}{
		{
			name:   "paypal is returned",
			expect: "paypal",
			input: mollie.Payment{
				Method: mollie.PayPal,
			},
		},
		{
			name:   "none is returned",
			expect: "none",
			input:  mollie.Payment{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := getSafePaymentMethod(c.input)
			assert.Equal(t, c.expect, got)
		})
	}
}
