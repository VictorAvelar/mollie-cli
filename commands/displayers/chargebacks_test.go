package displayers

import (
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMollieChargeback_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	disp := MollieChargeback{
		&mollie.Chargeback{
			ID:               "chg_test",
			PaymentID:        "tr_test",
			Amount:           &mollie.Amount{Currency: "USD", Value: "10.00"},
			SettlementAmount: &mollie.Amount{Currency: "USD", Value: "12.00"},
			CreatedAt:        &n,
		},
	}

	out := expectedChargebackSlice(*disp.Chargeback)
	assert.Len(t, out, 1)
	assert.Equal(t, out, disp.KV())
}

func TestMollieListChargebacks(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	var cbs []mollie.Chargeback
	{
		cbs = append(cbs, mollie.Chargeback{
			ID:               "chg_test",
			PaymentID:        "tr_test",
			Amount:           &mollie.Amount{Currency: "USD", Value: "10.00"},
			SettlementAmount: &mollie.Amount{Currency: "USD", Value: "12.00"},
			CreatedAt:        &n,
		},
			mollie.Chargeback{
				ID:               "chg_test_2",
				PaymentID:        "tr_test_2",
				Amount:           &mollie.Amount{Currency: "USD", Value: "100.00"},
				SettlementAmount: &mollie.Amount{Currency: "USD", Value: "120.00"},
				CreatedAt:        &n,
			},
		)
	}

	disp := MollieChargebackList{
		ChargebackList: &mollie.ChargebackList{
			Count: 2,
			Embedded: struct{ Chargebacks []mollie.Chargeback }{
				Chargebacks: cbs,
			},
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	out := expectedChargebackSlice(disp.Embedded.Chargebacks...)
	assert.Len(t, out, 2)
	assert.Equal(t, disp.ChargebackList.Count, 2)
	assert.Equal(t, out, disp.KV())
}

func expectedChargebackSlice(cbs ...mollie.Chargeback) (out []map[string]interface{}) {
	for _, c := range cbs {
		x := map[string]interface{}{
			"ID":         c.ID,
			"Payment":    c.PaymentID,
			"Amount":     stringCombinator(" ", c.Amount.Value, c.Amount.Currency),
			"Settlement": stringCombinator(" ", c.SettlementAmount.Value, c.SettlementAmount.Currency),
			"Created at": c.CreatedAt.Format("02-01-2006"),
		}

		out = append(out, x)
	}

	return
}
