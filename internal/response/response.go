package response

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseError(message string, code int) *Response[string] {
	return &Response[string]{
		Code:    code,
		Message: message,
		Data:    "",
	}
}

func ResponseSuccess[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func ResponseCreateSuccess() *Response[string] {
	return &Response[string]{
		Code:    201,
		Message: "created successfully",
		Data:    "",
	}
}
