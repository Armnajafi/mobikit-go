package rules

type RequiredRule struct {
	Message string
}

func (r RequiredRule) Validate(value interface{}) (bool, string) {
	if value == nil || value == "" {
		return false, r.Message
	}
	return true, ""
}
