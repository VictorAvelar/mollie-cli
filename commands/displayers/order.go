package displayers

import "github.com/VictorAvelar/mollie-api-go/v3/mollie"

// MollieOrderList wrapper for displaying.
type MollieOrderList struct {
	*mollie.OrderList
}

// KV is a displayable group of key value.
func (mpl *MollieOrderList) KV() []map[string]interface{} {
	out := outPrealloc(len(mpl.Embedded.Orders))

	for _, r := range mpl.Embedded.Orders {
		x := buildXOrder(r)

		out = append(out, x)
	}

	return out
}

// MollieOrder wrapper for displaying.
type MollieOrder struct {
	*mollie.Order
}

// KV is a displayable group of key value.
func (mp *MollieOrder) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXOrder(mp.Order)
	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (mp *MollieOrder) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"OrderNumber",
		"Description",
		"ShippingAddress",
		"Amount",
		"PaidAt",
		"Status",
		"CREATED_AT",
	}
}

// ColMap returns a list of columns and its description.
func (mp *MollieOrder) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":        "the resource name",
		"ID":              "the order id",
		"OrderNumber":     "the order number",
		"Description":     "the Description",
		"ShippingAddress": "the shipping address",
		"Amount":          "the order amount",
		"PaidAt":          "the order paid at",
		"Status":          "the order status",
		"CREATED_AT":      "the order creation date",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mp *MollieOrder) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mp *MollieOrder) Filterable() bool {
	return true
}

func buildXOrder(p *mollie.Order) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":        p.Resource,
		"ID":              p.ID,
		"OrderNumber":     p.OrderNumber,
		"Description":     p.Description,
		"ShippingAddress": p.ShippingAddress,
		"Amount":          p.Amount,
		"PaidAt":          p.PaidAt,
		"Status":          p.Status,
		"CREATED_AT":      p.CreatedAt,
	}
}
