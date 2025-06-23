package rules

type UppercaseRule struct {
	Message string
}

func (r UppercaseRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		for _, char := range v {
			if char < 'A' || char > 'Z' {
				return false, r.Message
			}
		}
	}
	return true, ""
}
