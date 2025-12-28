package validation

import "testing"

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func TestValidatorValidate_Success(t *testing.T) {
	v := New()

	input := LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	errs := v.Validate(input)

	if errs != nil {
		t.Errorf("expected no errors, got %v", errs)
	}
}
func TestValidatorValidate_Failure(t *testing.T) {
	v := New()

	input := LoginRequest{}

	errs := v.Validate(input)

	if errs == nil {
		t.Fatal("expected validation errors")
	}

	if errs["email"] == "" {
		t.Error("expected email error")
	}

	if errs["password"] == "" {
		t.Error("expected password error")
	}
}

