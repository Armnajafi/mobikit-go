package rules

type MaxLengthRule struct {
	Max     int
	Message string
}

func (r MaxLengthRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		if len(v) < r.Max {
			return false, r.Message
		}
	}
	return true, ""
}
