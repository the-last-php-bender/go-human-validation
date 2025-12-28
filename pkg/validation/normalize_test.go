package validation

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "Email", "email"},
		{"two words", "FirstName", "first_name"},
		{"three words", "UserAccountID", "user_account_i_d"},
		{"already lowercase", "password", "password"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
