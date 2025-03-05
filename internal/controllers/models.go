package controller



type ItemResult[T any] struct {
	Items []T `json:"items"`
}

type ErrorSchema struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    map[string]struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"data"`
}




type SuccessListResponse[T any] struct {
	Result ItemResult[T] `json:"result"`
	Error  any `json:"error"`
}

type ErrorListResponse struct {
	Result any `json:"result"`
	Error  ErrorSchema `json:"error"`
}

