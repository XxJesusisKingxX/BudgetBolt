package request

import (
	"encoding/json"
)

func ParseResponse(resp []byte, data interface{}) error {
	err := json.Unmarshal(resp, data)
	if err != nil {
		return err
	}
	return err
}