package rules

type RequiredIfRule struct {
	ConditionValue interface{}
	Message        string
}

func (r RequiredIfRule) Validate(value interface{}) (bool, string) {
	if v, ok := r.ConditionValue.(string); ok {
		if v != "" {
			if value == nil || value == "" {
				return false, r.Message
			}
		}
	}
	return true, ""
}
