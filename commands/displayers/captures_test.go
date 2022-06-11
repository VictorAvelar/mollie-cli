package displayers

import (
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/stretchr/testify/assert"
)

func TestCapturesKV(t *testing.T) {
	mc := &MollieCapture{
		Capture: &mollie.Capture{
			Resource: "captures",
			ID:       "cp_test",
		},
	}

	expect := []map[string]interface{}{{"AMOUNT": "--- ---", "CREATED_AT": "----------", "ID": "cp_test", "MODE": "", "PAYMENT_ID": "", "RESOURCE": "captures", "SETTLEMENT_AMOUNT": "--- ---", "SETTLEMENT_ID": "", "SHIPMENT_ID": ""}}
	assert.Equal(t, expect, mc.KV())
}

func TestCapturesListKV(t *testing.T) {
	var cl []*mollie.Capture

	cl = append(cl, &mollie.Capture{
		Resource: "captures",
		ID:       "cp_test",
	})
	mcl := MollieCapturesList{
		CapturesList: &mollie.CapturesList{
			Count: 1,
			Embedded: struct{ Captures []*mollie.Capture }{
				Captures: cl,
			},
			Links: mollie.PaginationLinks{
				Documentation: &mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          &mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	w := map[string]interface{}{"AMOUNT": "--- ---", "CREATED_AT": "----------", "ID": "cp_test", "MODE": "", "PAYMENT_ID": "", "RESOURCE": "captures", "SETTLEMENT_AMOUNT": "--- ---", "SETTLEMENT_ID": "", "SHIPMENT_ID": ""}

	var want []map[string]interface{}
	want = append(want, w)

	assert.Len(t, want, mcl.Count)
	assert.Equal(t, want, mcl.KV())
}
