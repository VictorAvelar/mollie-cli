package displayers

import (
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMollieRefund_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	disp := MollieRefund{
		Refund: &mollie.Refund{
			ID:           "rf_test",
			PaymentID:    "tr_test",
			OrderID:      "or_test",
			SettlementID: "stl_test",
			Amount: &mollie.Amount{
				Value:    "10.00",
				Currency: "EUR",
			},
			Status:      "paid",
			Description: "a test description",
			CreatedAt:   &n,
		},
	}

	out := expectedRefundSlice(disp.Refund)
	assert.Len(t, out, 1)
	assert.Equal(t, out, disp.KV())
}

func TestMollieRefundList_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	tomorrow := n.AddDate(0, 0, 1)
	var refunds []*mollie.Refund
	{
		refunds = append(refunds, &mollie.Refund{
			ID:           "rf_test",
			PaymentID:    "tr_test",
			OrderID:      "or_test",
			SettlementID: "stl_test",
			Amount: &mollie.Amount{
				Value:    "10.00",
				Currency: "EUR",
			},
			Status:      "paid",
			Description: "a test description",
			CreatedAt:   &n,
		}, &mollie.Refund{
			ID:           "rf_test_alt",
			PaymentID:    "tr_test_alt",
			OrderID:      "or_test_alt",
			SettlementID: "stl_test_alt",
			Amount: &mollie.Amount{
				Value:    "100.00",
				Currency: "EUR",
			},
			Status:      "queued",
			Description: "a test_alt description",
			CreatedAt:   &tomorrow,
		})
	}

	disp := MollieRefundList{
		RefundList: &mollie.RefundList{
			Count: 2,
			Embedded: struct{ Refunds []*mollie.Refund }{
				Refunds: refunds,
			},
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	out := expectedRefundSlice(disp.Embedded.Refunds...)

	assert.Len(t, out, 2)
	assert.Equal(t, 2, disp.Count)
	assert.Equal(t, out, disp.KV())
}

func expectedRefundSlice(refs ...*mollie.Refund) (out []map[string]interface{}) {
	for _, r := range refs {
		x := map[string]interface{}{
			"RESOURCE":          r.Resource,
			"ID":                r.ID,
			"AMOUNT":            fallbackSafeAmount(r.Amount),
			"SETTLEMENT_ID":     r.SettlementID,
			"SETTLEMENT_AMOUNT": fallbackSafeAmount(r.SettlementAmount),
			"DESCRIPTION":       r.Description,
			"METADATA":          r.Metadata,
			"STATUS":            r.Status,
			"PAYMENT_ID":        r.PaymentID,
			"ORDER_ID":          r.OrderID,
			"CREATED_AT":        fallbackSafeDate(r.CreatedAt),
		}

		out = append(out, x)
	}

	return out
}
