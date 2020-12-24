package displayers

import (
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMollieCustomers_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "24-12-2020")
	if err != nil {
		t.Error(err)
	}
	disp := MollieCustomer{
		Customer: &mollie.Customer{
			Resource:  "customer",
			ID:        "cs_test",
			Name:      "test customer",
			Email:     "test@example.com",
			Locale:    mollie.German,
			CreatedAt: &n,
			Mode:      "test",
		},
	}

	want := []map[string]interface{}{
		{
			"CREATED_AT": "24-12-2020",
			"EMAIL":      "test@example.com",
			"ID":         "cs_test",
			"LOCALE":     "de_DE",
			"METADATA":   map[string]interface{}(nil),
			"MODE":       "test",
			"NAME":       "test customer",
			"RESOURCE":   "customer",
		},
	}

	assert.Len(t, want, 1)
	assert.Equal(t, want, disp.KV())
}

func TestMollieCustomersList_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "24-12-2020")
	if err != nil {
		t.Error(err)
	}

	var customers []mollie.Customer
	{
		customers = append(
			customers,
			mollie.Customer{
				Resource:  "customer",
				ID:        "cs_test",
				Name:      "test customer",
				Email:     "test@example.com",
				Locale:    mollie.German,
				CreatedAt: &n,
				Mode:      "test",
			},
			mollie.Customer{
				Resource:  "customer",
				ID:        "cs_test_2",
				Name:      "test customer 2",
				Email:     "test2@example.com",
				Locale:    mollie.Spanish,
				CreatedAt: &n,
				Mode:      "test",
			},
		)
	}

	disp := MollieCustomerList{
		CustomersList: &mollie.CustomersList{
			Count: 2,
			Embedded: struct {
				Customers []mollie.Customer "json:\"customers,omitempty\""
			}{
				Customers: customers,
			},
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	want := []map[string]interface{}{
		{
			"CREATED_AT": "24-12-2020",
			"EMAIL":      "test@example.com",
			"ID":         "cs_test",
			"LOCALE":     "de_DE",
			"METADATA":   map[string]interface{}(nil),
			"MODE":       "test",
			"NAME":       "test customer",
			"RESOURCE":   "customer",
		},
		{
			"CREATED_AT": "24-12-2020",
			"EMAIL":      "test2@example.com",
			"ID":         "cs_test_2",
			"LOCALE":     "es_ES",
			"METADATA":   map[string]interface{}(nil),
			"MODE":       "test",
			"NAME":       "test customer 2",
			"RESOURCE":   "customer",
		},
	}

	assert.Len(t, want, disp.Count)
	assert.Equal(t, want, disp.KV())
}
