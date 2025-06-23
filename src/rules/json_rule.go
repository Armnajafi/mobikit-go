package rules

import (
	"encoding/json"
)

type JSONRule struct {
	Message string
}

func (r JSONRule) Validate(value interface{}) (bool, string) {
	if v, ok := value.(string); ok {
		var js json.RawMessage
		if err := json.Unmarshal([]byte(v), &js); err != nil {
			return false, r.Message
		}
	}
	return true, ""
}
