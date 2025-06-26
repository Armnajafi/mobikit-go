package rules

type MaxRule struct {
	Max     int
	Message string
}

func (r MaxRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(int); ok {
		if v > r.Max {
			return false, r.Message
		}
	}
	return true, ""
}
