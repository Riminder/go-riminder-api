package response

// Container is the base representation of a response.
type Container struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
