package domain

import "net/http"

type ErrorResponse struct {
	Message string `json:"message"`
}

type ResponseDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ErrorException struct {
	Response   ResponseDto
	StatusCode int
}

func NewErrorException(response ResponseDto, statusCode int) *ErrorException {
	return &ErrorException{
		Response:   response,
		StatusCode: statusCode,
	}
}

// Keyspace error response
func BadRequest() *ErrorException {
	return NewErrorException(ResponseDto{
		Code:    "KS001",
		Message: "Bad Request",
		Success: false,
	}, http.StatusBadRequest)
}

func BadRequestWith(message string) *ErrorException {
	if message == "" {
		message = "Bad Request"
	}

	return NewErrorException(ResponseDto{
		Code:    "KS001",
		Message: message,
		Success: false,
	}, http.StatusBadRequest)
}

func UnAuthorized() *ErrorException {
	return NewErrorException(ResponseDto{
		Code:    "KS002",
		Message: "UnAuthorized",
		Success: false,
	}, http.StatusUnauthorized)
}

func Forbidden() *ErrorException {
	return NewErrorException(ResponseDto{
		Code:    "KS003",
		Message: "Forbidden",
		Success: false,
	}, http.StatusForbidden)
}
