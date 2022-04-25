package response

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func GetErrorResponse(statusCode int, message string) errorResponse {
	return errorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
