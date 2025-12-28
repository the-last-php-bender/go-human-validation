package validation

import "testing"

func TestDefaultMessages(t *testing.T) {
	msgs := defaultMessages()

	if msgs["required"] == "" {
		t.Error("expected required message to exist")
	}

	if msgs["password.min"] == "" {
		t.Error("expected password.min message to exist")
	}
}
