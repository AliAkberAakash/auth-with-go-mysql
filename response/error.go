package response

type errorResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func GetCommonResponse(statusCode int, message string, data interface{}) errorResponse {
	return errorResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
