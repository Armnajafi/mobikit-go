package rules

type MinRule struct {
	Min     int
	Message string
}

func (r MinRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(int); ok {
		if v < r.Min {
			return false, r.Message
		}
	}
	return true, ""
}
