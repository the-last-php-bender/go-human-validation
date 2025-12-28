package validation

import "testing"

func TestWithMessagesOption(t *testing.T) {
	custom := map[string]string{
		"required": "Custom required message",
	}

	v := New(WithMessages(custom))

	if v.messages["required"] != "Custom required message" {
		t.Error("expected custom message to override default")
	}
}
