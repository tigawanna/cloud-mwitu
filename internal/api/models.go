package api

type ListResponse[T any] struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
	Items   T `json:"items"`
}


