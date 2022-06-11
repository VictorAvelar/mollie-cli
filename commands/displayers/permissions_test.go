package displayers

import (
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMolliePermission_KV(t *testing.T) {
	perm := mollie.Permission{
		Description: "random desc",
		Granted:     true,
		ID:          "random.test",
		Resource:    "test_resource",
	}

	disp := MolliePermission{
		Permission: &perm,
	}

	out := []map[string]interface{}{}
	out = append(out, buildXPermission(&perm))

	assert.Len(t, disp.KV(), 1)
	assert.Equal(t, out, disp.KV())
}

func TestMolliePermissionList_KV(t *testing.T) {
	perm := mollie.Permission{
		Description: "random desc",
		Granted:     true,
		ID:          "random.test",
		Resource:    "test_resource",
	}
	permList := MolliePermissionList{
		PermissionsList: &mollie.PermissionsList{
			Count: 2,
			Embedded: struct {
				Permissions []*mollie.Permission "json:\"permissions,omitempty\""
			}{
				[]*mollie.Permission{&perm, &perm},
			},
			Links: mollie.PermissionLinks{
				Documentation: &mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          &mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	out := []map[string]interface{}{}
	out = append(out, buildXPermission(&perm), buildXPermission(&perm))

	assert.Len(t, permList.KV(), 2)
	assert.Equal(t, out, permList.KV())

}
