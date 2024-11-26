package response

// ResponseErr is a struct for error response
type ResponseErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseErrorNew(message string, code int) *ResponseErr {
	return &ResponseErr{
		Code:    code,
		Message: message,
	}
}

func ResponseErrorConflict(errMess string) *ResponseErr {
	return &ResponseErr{
		Code: 409,
		// Message: "data already exist",
		Message: errMess,
	}
}

func ResponseErrorNotFound() *ResponseErr {
	return &ResponseErr{
		Code:    404,
		Message: "data not found",
	}
}

func ResponseErrorInternalServer() *ResponseErr {
	return &ResponseErr{
		Code:    500,
		Message: "internal server error",
	}
}

func ResponseErrorBadRequest() *ResponseErr {
	return &ResponseErr{
		Code:    400,
		Message: "bad request",
	}
}

func ResponseErrorUnauthorized() *ResponseErr {
	return &ResponseErr{
		Code:    401,
		Message: "unauthorized",
	}
}

func ResponseErrorForbidden() *ResponseErr {
	return &ResponseErr{
		Code:    403,
		Message: "forbidden",
	}
}

func ResponseErrorUnprocessableEntity() *ResponseErr {
	return &ResponseErr{
		Code:    422,
		Message: "unprocessable entity",
	}
}

func ResponseError(message string, code int) *Response[string] {
	return &Response[string]{
		Code:    code,
		Message: message,
		Data:    "",
	}
}
