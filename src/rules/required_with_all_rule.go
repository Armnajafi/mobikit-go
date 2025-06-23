package rules

type RequiredWithAllRule struct {
	OtherValues []interface{}
	Message     string
}

func (r RequiredWithAllRule) Validate(value interface{}) (bool, string) {
	allPresent := true
	for _, otherValue := range r.OtherValues {
		if otherValue == nil || otherValue == "" {
			allPresent = false
			break
		}
	}

	if allPresent {
		if value == nil || value == "" {
			return false, r.Message
		}
	}

	return true, ""
}
