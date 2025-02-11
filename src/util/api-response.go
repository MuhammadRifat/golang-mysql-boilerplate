package util

import "net/http"

type ResponseType[T any] struct {
	Pagination *Pagination `json:"Pagination,omitempty"`
	Results    T           `json:"Results"`
}

func ApiResponse[T any](StatusCode int, Data T, Pagination ...Pagination) (int, ResponseType[T]) {
	if len(Pagination) > 0 {
		return StatusCode, ResponseType[T]{Pagination: &Pagination[0], Results: Data}
	}
	return StatusCode, ResponseType[T]{Results: Data}
}

func ResponseOK[T any](Data T, Pagination ...Pagination) (int, ResponseType[T]) {
	return ApiResponse(http.StatusOK, Data, Pagination...)
}

func ResponseCreated[T any](Data T) (int, ResponseType[T]) {
	return ApiResponse(int(http.StatusCreated), Data)
}
