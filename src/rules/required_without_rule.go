package rules

type RequiredWithoutRule struct {
	OtherValue interface{}
	Message    string
}

func (r RequiredWithoutRule) Validate(value interface{}) (bool, string) {
	if r.OtherValue == nil || r.OtherValue == "" {
		if value == nil || value == "" {
			return false, r.Message
		}
	}
	return true, ""
}
