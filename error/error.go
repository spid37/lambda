package error

import (
	"encoding/json"
	"log"
)

// Error -
type Error struct {
	Code    string
	Message string

	OrigErr error
}

func (e *Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		log.Println("cannot marshal Error:", e)
		panic(err)
	}
	return string(b[:])
}

// MarshalJSON outputs error in json
func (e *Error) MarshalJSON() ([]byte, error) {
	var origErrMsg string
	if e.OrigErr != nil {
		origErrMsg = e.OrigErr.Error()
	}
	return json.Marshal(&struct {
		Code           string `json:"code"`
		PublicMessage  string `json:"public_message"`
		PrivateMessage string `json:"private_message"`
	}{
		Code:           e.Code,
		PublicMessage:  e.Message,
		PrivateMessage: origErrMsg,
	})
}
