package errors

import (
	"fmt"
)

// ResponseError is an error that occurs when a request status code was > 200.
type ResponseError struct {
	code    int
	message string
	body    string
}

// NewResponseError creates a response error.
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

// Error return error string.
func (r *ResponseError) Error() string {
	return fmt.Sprintf("Invalid Response: %d -> %s (%s)", r.code, r.message, r.body)
}
