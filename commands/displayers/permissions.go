package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MolliePermissionList is wrapper for displaying.
type MolliePermissionList struct {
	*mollie.PermissionsList
}

// KV is a displayable group of key value.
func (mp *MolliePermissionList) KV() []map[string]interface{} {
	out := outPrealloc(len(mp.Embedded.Permissions))

	for _, p := range mp.Embedded.Permissions {
		x := buildXPermission(p)

		out = append(out, x)
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (mp *MolliePermissionList) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"GRANTED",
	}
}

// ColMap returns a list of columns and its description.
func (mp *MolliePermissionList) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":    "the resource name specified by mollie",
		"ID":          "the permissions id",
		"DESCRIPTION": "the permission description",
		"GRANTED":     "the permission status for the current api/access token",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mp *MolliePermissionList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mp *MolliePermissionList) Filterable() bool {
	return true
}

// MolliePermission is wrapper for displaying.
type MolliePermission struct {
	*mollie.Permission
}

// KV is a displayable group of key value.
func (p *MolliePermission) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXPermission(p.Permission)
	out = append(out, x)
	return out
}

// Cols returns an array of columns available for displaying.
func (p *MolliePermission) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"GRANTED",
	}
}

// ColMap returns a list of columns and its description.
func (p *MolliePermission) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":    "the resource name specified by mollie",
		"ID":          "the permissions id",
		"DESCRIPTION": "the permission description",
		"GRANTED":     "the permission status for the current api/access token",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (p *MolliePermission) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (p *MolliePermission) Filterable() bool {
	return true
}

func buildXPermission(mp *mollie.Permission) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":    mp.Resource,
		"ID":          mp.ID,
		"DESCRIPTION": mp.Description,
		"GRANTED":     mp.Granted,
	}
}
