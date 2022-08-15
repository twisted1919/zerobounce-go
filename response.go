package zerobounce

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseType[T CreditsResponseSuccess | ValidateResponseSuccess] struct {
	Success *T
	Error   *ResponseError
}

func (r *ResponseType[T]) IsSuccess() bool {
	return r.Success != nil && r.Error == nil
}

func (r *ResponseType[T]) IsError() bool {
	return r.Error != nil && r.Success == nil
}

func newResponseSuccess[T CreditsResponseSuccess | ValidateResponseSuccess](response *T) *ResponseType[T] {
	return &ResponseType[T]{
		Success: response,
		Error:   nil,
	}
}

func newResponseError[T CreditsResponseSuccess | ValidateResponseSuccess](response *ResponseError) *ResponseType[T] {
	return &ResponseType[T]{
		Success: nil,
		Error:   response,
	}
}
