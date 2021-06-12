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

// Cols returns an array of columns available for displaying.
func (mc *MollieCustomer) Cols() []string {
	return customersCols()
}

// ColMap returns a list of columns and its description.
func (mc *MollieCustomer) ColMap() map[string]string {
	return customersColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mc *MollieCustomer) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mc *MollieCustomer) Filterable() bool {
	return true
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

// Cols returns an array of columns available for displaying.
func (mcl *MollieCustomerList) Cols() []string {
	return customersCols()
}

// ColMap returns a list of columns and its description.
func (mcl *MollieCustomerList) ColMap() map[string]string {
	return customersColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mcl *MollieCustomerList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mcl *MollieCustomerList) Filterable() bool {
	return true
}

func customersCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"EMAIL",
		"LOCALE",
		"METADATA",
		"CREATED_AT",
	}
}

func customersColMap() map[string]string {
	return map[string]string{
		"RESOURCE":   "the resource name",
		"ID":         "the resource id",
		"MODE":       "the mode used to create this customer.",
		"NAME":       "the full name of the customer",
		"EMAIL":      "the email address of the customer",
		"LOCALE":     "language to be used in the hosted payment pages shown to the consumer",
		"METADATA":   "any data you like to attach to the customer",
		"CREATED_AT": "the customer’s date and time of creation",
	}
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
