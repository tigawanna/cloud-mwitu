package models

type BareSuccessResponse[Res any] struct {
	StatusCode int    `json:"statusCode"`
	Result     Res    `json:"result"`
	Message    string `json:"message"`
}
