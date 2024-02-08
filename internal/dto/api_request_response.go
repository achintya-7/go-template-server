package dto

// ApiResponse is a standard response from the service.
type ApiResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  *ApiError   `json:"error"`
}

// ApiError is a standard error response from the service.
type ApiError struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int
}

// ErrorResponse is a standard error response from handler methods.
type ErrorResponse struct {
	Code           int
	Message        string
	HttpStatusCode int
}