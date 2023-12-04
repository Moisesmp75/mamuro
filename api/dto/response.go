package dto

import (
	"encoding/json"
)

type BaseResponse[T any] struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Resource T      `json:"resource"`
}

func NewResponse[T any](resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Resource: resource,
		Message:  "",
	}
}

func ErrorResponse(errorMessage string) BaseResponse[*string] {
	return BaseResponse[*string]{
		Success:  false,
		Message:  errorMessage,
		Resource: nil,
	}
}

func JsonResponse[T any](response any) string {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return "Error al convertir a json"
	}
	return string(jsonData)
}