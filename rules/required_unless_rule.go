package rules

type RequiredUnlessRule struct {
	UnlessValue interface{}
	Message     string
}

func (r RequiredUnlessRule) Validate(value interface{}) (bool, string) {
	if v, ok := r.UnlessValue.(string); ok {
		if v == "" {
			if value == nil || value == "" {
				return false, r.Message
			}
		}
	}
	return true, ""
}
