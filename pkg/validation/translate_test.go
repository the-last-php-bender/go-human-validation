package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type testInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func TestTranslateValidationErrors(t *testing.T) {
	v := New()

	input := testInput{}
	err := v.validate.Struct(input)

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Fatal("expected ValidationErrors")
	}

	result := v.translate(ve)

	if result["email"] == "" {
		t.Error("expected email error message")
	}

	if result["password"] == "" {
		t.Error("expected password error message")
	}
}
