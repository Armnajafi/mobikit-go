package rules

import (
	"regexp"
)

type RegexRule struct {
	Pattern string
	Message string
}

func (r RegexRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		re := regexp.MustCompile(r.Pattern)
		if !re.MatchString(v) {
			return false, r.Message
		}
	}
	return true, ""
}
