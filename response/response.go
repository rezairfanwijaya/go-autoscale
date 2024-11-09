package response

import (
	"encoding/json"
	"log"
)

type SuccessResp struct {
	Data       any `json:"data"`
	StatusCode int `json:"status_code"`
}

func (s *SuccessResp) ChangToByte() ([]byte, error) {
	res, err := json.Marshal(s)
	if err != nil {
		log.Println("failed change response to bye")
		return res, err
	}

	return res, nil
}
