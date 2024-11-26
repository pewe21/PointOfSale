package response

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
