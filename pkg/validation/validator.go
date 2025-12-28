package validation

import (
	"github.com/go-playground/validator/v10"
)
// Validator is the main struct exposed by this package.
// It wraps go-playground/validator and adds human-readable error handling.
type Validator struct {
	validate *validator.Validate
	messages map[string]string
}
// New creates a new Validator instance.
// It accepts functional options to allow future extensibility
// without breaking the public API.
func New(opts ...Option) *Validator {
	v := &Validator{
		validate: validator.New(),
		messages: defaultMessages(),
	}

	for _, opt := range opts {
		opt(v)
	}

	return v
}
// Validate validates a struct and returns a map of human-readable errors.
// If validation passes, it returns nil.
func (v *Validator) Validate(input any) map[string]string {
	err := v.validate.Struct(input)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"_error": "Invalid input",
		}
	}

	return v.translate(validationErrors)
}
