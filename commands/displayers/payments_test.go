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
	tomorrow := n.AddDate(0, 0, 2)
	disp := MolliePayment{
		Payment: &mollie.Payment{
			ID:            "tr_test",
			Mode:          mollie.TestMode,
			Status:        "paid",
			CreatedAt:     &n,
			ExpiresAt:     &tomorrow,
			IsCancellable: false,
			Amount:        &mollie.Amount{Currency: "EUR", Value: "1.00"},
			Method:        mollie.PayPal,
			Description:   "testing KV",
			Locale:        mollie.Locale(""),
		},
	}

	w := map[string]interface{}{
		"AMOUNT":          "€1.00",
		"APP_FEE":         "none",
		"AUTHORIZED_AT":   "----------",
		"CANCELABLE":      false,
		"CANCELED_AT":     "----------",
		"CAPTURED":        "--- ---",
		"COUNTRY":         "",
		"CREATED_AT":      "04-11-2020",
		"CUSTOMER_ID":     "",
		"DESCRIPTION":     "testing KV",
		"EXPIRES":         "06-11-2020",
		"FAILED_AT":       "----------",
		"ID":              "tr_test",
		"LOCALE":          "",
		"MANDATE_ID":      "",
		"METHOD":          "paypal",
		"MODE":            "test",
		"ORDER_ID":        "",
		"PAID_AT":         "----------",
		"REDIRECT":        "",
		"REFUNDED":        "--- ---",
		"REMAINING":       "--- ---",
		"RESOURCE":        "",
		"SEQUENCE":        "",
		"SETTLEMENT":      "--- ---",
		"SETTLEMENT_ID":   "",
		"STATUS":          "paid",
		"SUBSCRIPTION_ID": "",
		"WEBHOOK":         "",
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
	tomorrow := n.AddDate(0, 0, 2)
	var ps []mollie.Payment
	{
		ps = append(
			ps,
			mollie.Payment{
				ID:            "tr_test",
				Mode:          mollie.TestMode,
				Status:        "paid",
				CreatedAt:     &n,
				ExpiresAt:     &tomorrow,
				IsCancellable: false,
				Amount:        &mollie.Amount{Currency: "EUR", Value: "1.00"},
				Method:        mollie.PayPal,
				Description:   "testing KV",
			},
			mollie.Payment{
				ID:            "tr_test_2",
				Mode:          mollie.TestMode,
				Status:        "expired",
				CreatedAt:     &n,
				ExpiresAt:     &tomorrow,
				IsCancellable: false,
				Amount:        &mollie.Amount{Currency: "USD", Value: "2.00"},
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

	w := map[string]interface{}{"AMOUNT": "€1.00", "APP_FEE": "none", "AUTHORIZED_AT": "----------", "CANCELABLE": false, "CANCELED_AT": "----------", "CAPTURED": "--- ---", "COUNTRY": "", "CREATED_AT": "04-11-2020", "CUSTOMER_ID": "", "DESCRIPTION": "testing KV", "EXPIRES": "06-11-2020", "FAILED_AT": "----------", "ID": "tr_test", "LOCALE": "", "MANDATE_ID": "", "METHOD": "paypal", "MODE": "test", "ORDER_ID": "", "PAID_AT": "----------", "REDIRECT": "", "REFUNDED": "--- ---", "REMAINING": "--- ---", "RESOURCE": "", "SEQUENCE": "", "SETTLEMENT": "--- ---", "SETTLEMENT_ID": "", "STATUS": "paid", "SUBSCRIPTION_ID": "", "WEBHOOK": ""}

	w1 := map[string]interface{}{"AMOUNT": "$2.00", "APP_FEE": "none", "AUTHORIZED_AT": "----------", "CANCELABLE": false, "CANCELED_AT": "----------", "CAPTURED": "--- ---", "COUNTRY": "", "CREATED_AT": "04-11-2020", "CUSTOMER_ID": "", "DESCRIPTION": "testing KV list payments", "EXPIRES": "06-11-2020", "FAILED_AT": "----------", "ID": "tr_test_2", "LOCALE": "", "MANDATE_ID": "", "METHOD": "banktransfer", "MODE": "test", "ORDER_ID": "", "PAID_AT": "----------", "REDIRECT": "", "REFUNDED": "--- ---", "REMAINING": "--- ---", "RESOURCE": "", "SEQUENCE": "", "SETTLEMENT": "--- ---", "SETTLEMENT_ID": "", "STATUS": "expired", "SUBSCRIPTION_ID": "", "WEBHOOK": ""}

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
				ExpiresAt: &n,
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
