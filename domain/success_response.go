package domain

type Meta struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
	Total  uint `json:"total"`
}

type ApiResponse[T any] struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Meta    *Meta  `json:"meta,omitempty"`
	Data    *T     `json:"data,omitempty"`
}

func NewApiResponse[T any](code string, success bool, message string, meta *Meta, data *T) *ApiResponse[T] {
	return &ApiResponse[T]{
		Code:    code,
		Success: success,
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}

func OkApiResponse[T any](data *T, message string, meta *Meta) *ApiResponse[T] {
	if message == "" {
		message = "Success"
	}

	return NewApiResponse("KS000", true, message, meta, data)
}
