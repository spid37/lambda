package error

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestLambdaError(t *testing.T) {
	expectedString := `{"code":"SECRET_MESSAGE","public_message":"Failed to complete","private_message":"This is the original error"}`

	is := is.New(t)
	newerr := &Error{
		Code:    "SECRET_MESSAGE",
		Message: "Failed to complete",
		OrigErr: errors.New("This is the original error"),
	}

	b, err := json.Marshal(newerr)
	is.NoErr(err)
	is.Equal(expectedString, string(b))
}
