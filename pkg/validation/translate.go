package validation

import "github.com/go-playground/validator/v10"
func (v *Validator) translate(errs validator.ValidationErrors) map[string]string {
	result := make(map[string]string)

	for _, err := range errs {
		field := toSnakeCase(err.Field())
		tag := err.Tag()

		// Try field-specific message first
		key := field + "." + tag

		msg, ok := v.messages[key]
		if !ok {
			msg = v.messages[tag]
		}

		if msg == "" {
			msg = "Invalid value"
		}

		result[field] = msg
	}

	return result
}
