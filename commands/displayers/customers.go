package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieCustomer wrapper for displaying.
type MollieCustomer struct {
	*mollie.Customer
}

// KV is a displayable group of key value.
func (mc *MollieCustomer) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXCustomer(mc.Customer)

	out = append(out, x)

	return out
}

// MollieCustomerList wrapper for displaying.
type MollieCustomerList struct {
	*mollie.CustomersList
}

// KV is a displayable group of key value.
func (mcl *MollieCustomerList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, c := range mcl.Embedded.Customers {
		x := buildXCustomer(&c)

		out = append(out, x)
	}

	return out
}

func buildXCustomer(c *mollie.Customer) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":   c.Resource,
		"ID":         c.ID,
		"MODE":       fallbackSafeMode(c.Mode),
		"NAME":       c.Name,
		"EMAIL":      c.Email,
		"LOCALE":     fallbackSafeLocale(c.Locale),
		"METADATA":   c.Metadata,
		"CREATED_AT": fallbackSafeDate(c.CreatedAt),
	}
}
