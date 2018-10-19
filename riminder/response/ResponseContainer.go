package riminderResponse

// Container is the base representation of a riminderResponse.
type Container struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
