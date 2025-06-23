package rules

type MinLengthRule struct {
	Min     int
	Message string
}

func (r MinLengthRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		if len(v) < r.Min {
			return false, r.Message
		}
	}
	return true, ""
}
