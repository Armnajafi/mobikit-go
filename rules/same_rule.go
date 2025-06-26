package rules

type SameRule struct {
	OtherValue interface{}
	Message    string
}

func (r SameRule) Validate(value interface{}) (bool, string) {
	if value != r.OtherValue {
		return false, r.Message
	}
	return true, ""
}
