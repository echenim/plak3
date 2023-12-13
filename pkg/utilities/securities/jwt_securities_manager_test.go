package securities

import "testing"

func TestIsAllowedRole(t *testing.T) {
	type testdata struct {
		name         string
		roles        []string
		allowedRoles string
		expected     bool
	}

	data := []testdata{
		{
			name:         "test admin",
			roles:        []string{"admin", "view", "edit", "delete"},
			allowedRoles: "admin",
			expected:     true,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			r := isAllowedRole(tt.roles, tt.allowedRoles)
			if !r {
				t.Errorf("expected %v, got %v", tt.expected, r)
			}
		})
	}
}
