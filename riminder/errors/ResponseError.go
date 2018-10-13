package errors

import (
	"fmt"
)

type ResponseError struct {
	code    int
	message string
	body    string
}

func NewResponseError(code int, message, body string) *ResponseError {
	r := new(ResponseError)
	r.code = code
	r.body = body
	r.message = message
	if len(r.body) > 200 {
		r.body = r.body[0:200]
	}
	return r
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("Invalid Response: %d -> %s (%s)", r.code, r.message, r.body)
}
