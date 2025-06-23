package rules

type RequiredWithoutAllRule struct {
	OtherValues []interface{}
	Message     string
}

func (r RequiredWithoutAllRule) Validate(value interface{}) (bool, string) {
	nonePresent := true
	for _, otherValue := range r.OtherValues {
		if otherValue != nil && otherValue != "" {
			nonePresent = false
			break
		}
	}

	if nonePresent {
		if value == nil || value == "" {
			return false, r.Message
		}
	}

	return true, ""
}
