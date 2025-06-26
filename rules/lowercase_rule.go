package rules

type LowercaseRule struct {
	Message string
}

func (r LowercaseRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		for _, char := range v {
			if char < 'a' || char > 'z' {
				return false, r.Message
			}
		}
	}
	return true, ""
}
