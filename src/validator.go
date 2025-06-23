package main

import (
	"encoding/json"
	"io/ioutil"
)

type Validator struct {
	locale string
}

type ValidationResult struct {
	IsValid bool
	Errors  map[string]string
}

type Rule interface {
	Validate(value interface{}) (bool, string)
}

func NewValidator() *Validator {
	return &Validator{
		locale: "en", // Default locale
	}
}

func (v *Validator) SetLocale(locale string) {
	v.locale = locale
}

func (v *Validator) Validate(inputs map[string]interface{}, rules map[string]Rule) ValidationResult {
	result := ValidationResult{
		IsValid: true,
		Errors:  make(map[string]string),
	}

	for field, value := range inputs {
		if rule, exists := rules[field]; exists {
			valid, message := rule.Validate(value)
			if !valid {
				result.IsValid = false
				result.Errors[field] = message
			}
		}
	}

	return result
}

// LoadLocale loads the error messages from the specified locale file
func (v *Validator) LoadLocale(locale string) error {
	filePath := "src/locales/" + locale + ".json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var messages map[string]string
	if err := json.Unmarshal(data, &messages); err != nil {
		return err
	}

	// Here you would store messages in the validator for use
	return nil
}
