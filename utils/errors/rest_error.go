package errors

import "net/http"

type RestErr struct {
	StatusCode int16  `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
		Error:      message,
	}
}