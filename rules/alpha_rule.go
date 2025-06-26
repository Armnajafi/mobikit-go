package rules

type AlphaRule struct {
	Message string
}

func (r AlphaRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		for _, char := range v {
			if !isAlpha(char) {
				return false, r.Message
			}
		}
	}
	return true, ""
}

func isAlpha(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
