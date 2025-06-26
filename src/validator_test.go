package src

import (
	"testing"

	"github.com/armnajafi/mobikit-go/src/rules"
)

func TestValidatorWithEmailRule(t *testing.T) {
	// ساخت validator
	validator := NewValidator()

	// ست کردن Rule برای ایمیل
	emailRule := rules.EmailRule{
		Message: "Invalid email address",
	}
	rulesMap := map[string]Rule{
		"email": emailRule,
	}

	// تست مقدار صحیح
	inputs := map[string]interface{}{
		"email": "test@example.com",
	}
	result := validator.Validate(inputs, rulesMap)
	if !result.IsValid {
		t.Errorf("Expected valid, got errors: %v", result.Errors)
	}

	// تست مقدار نادرست
	inputs["email"] = "invalid-email"
	result = validator.Validate(inputs, rulesMap)
	if result.IsValid {
		t.Errorf("Expected invalid, got valid")
	}
	if msg, ok := result.Errors["email"]; !ok || msg != "Invalid email address" {
		t.Errorf("Expected error message 'Invalid email address', got '%v'", msg)
	}
}
