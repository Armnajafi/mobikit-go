package rules

type NotInRule struct {
	Values  []string
	Message string
}

func (r NotInRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		for _, disallowed := range r.Values {
			if v == disallowed {
				return false, r.Message
			}
		}
	}
	return true, ""
}
