package rules

import (
	"regexp"
)

type EmailRule struct {
	Message string
}

func (r EmailRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !re.MatchString(v) {
			return false, r.Message
		}
	}
	return true, ""
}
