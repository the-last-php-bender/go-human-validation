package validation

import "unicode"
// toSnakeCase converts Go struct field names to snake_case.
// Example: FirstName → first_name
func toSnakeCase(input string) string {
	var out []rune

	for i, r := range input {
		if i > 0 && unicode.IsUpper(r) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(r))
	}

	return string(out)
}
