package helpers

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequest struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

type EmptyResponse struct{}

func BuildResponse(status string, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	// splitErr := strings.Split(err, "\n")

	res := Response{
		Status:  "Bad Request",
		Message: message,
		Data:    data,
	}

	return res
}

func BuildBadRequest(status string, message string, data struct{}) BadRequest {
	res := BadRequest{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}
