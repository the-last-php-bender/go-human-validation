package validation
// Option defines a functional option for configuring Validator.
type Option func(*Validator)
// WithMessages allows overriding or extending validation messages.
func WithMessages(messages map[string]string) Option {
	return func(v *Validator) {
		for key, msg := range messages {
			v.messages[key] = msg
		}
	}
}
