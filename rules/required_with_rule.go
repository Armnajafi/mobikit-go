package rules

type RequiredWithRule struct {
	OtherValue interface{}
	Message    string
}

func (r RequiredWithRule) Validate(value interface{}) (bool, string) {
	if r.OtherValue != nil && r.OtherValue != "" {
		if value == nil || value == "" {
			return false, r.Message
		}
	}
	return true, ""
}
