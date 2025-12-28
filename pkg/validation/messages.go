package validation
func defaultMessages() map[string]string {
	return map[string]string{
		// generic tags
		"required": "This field is required",
		"email":    "Must be a valid email address",
		"min":      "Value is too short",
		"max":      "Value is too long",
		"oneof":    "Invalid value",

		// field-specific overrides
		"password.required": "Password is required",
		"password.min":      "Password must be at least 8 characters",
	}
}
