package rules

type InRule struct {
	Values  []string
	Message string
}

func (r InRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		for _, allowed := range r.Values {
			if v == allowed {
				return true, ""
			}
		}
		return false, r.Message
	}
	return true, ""
}
