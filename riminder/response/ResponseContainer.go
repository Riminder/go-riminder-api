package response

type ResponseContainer struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
