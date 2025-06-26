package rules

import (
	"strconv"
)

type NumericRule struct {
	Message string
}

func (r NumericRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		_, err := strconv.Atoi(v)
		if err != nil {
			return false, r.Message
		}
	}
	return true, ""
}
