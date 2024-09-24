package utils

type TError struct {
	code int
	msg  string
	log  string
}

type TResponse[T struct{}] struct {
	Msg    string
	Data   T
	Errors []TError
}

func SuccessResponse[T struct{}](msg string, data T) TResponse[T] {
	return TResponse[T]{Msg: msg, Data: data, Error: nil}
}

func ErrorResponse(msg string, err []TError) TResponse[struct{}] {
	return TResponse[struct{}]{Msg: msg, Data: struct{}{}, Error: err}
}
