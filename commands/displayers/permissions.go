package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MolliePermissionList is wrapper for displaying.
type MolliePermissionList struct {
	*mollie.PermissionsList
}

// KV is a displayable group of key value
func (mp *MolliePermissionList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, p := range mp.Embedded.Permissions {
		x := buildXPermission(p)

		out = append(out, x)
	}

	return out
}

// MolliePermission is wrapper for displaying.
type MolliePermission struct {
	*mollie.Permission
}

// KV is a displayable group of key value
func (p *MolliePermission) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXPermission(p.Permission)
	out = append(out, x)
	return out
}

func buildXPermission(mp *mollie.Permission) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":    mp.Resource,
		"ID":          mp.ID,
		"DESCRIPTION": mp.Description,
		"GRANTED":     mp.Granted,
	}
}
