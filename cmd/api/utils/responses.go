package utils

type TResponse[T struct{}] struct {
	Msg   string
	Data  T
	Error []error
}

func SuccessResponse[T struct{}](msg string, data T) TResponse[T] {
	return TResponse[T]{Msg: msg, Data: data, Error: nil}
}

func ErrorResponse(msg string, err []error) TResponse[struct{}] {
	return TResponse[struct{}]{Msg: msg, Data: struct{}{}, Error: err}
}
