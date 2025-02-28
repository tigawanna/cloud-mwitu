package services

import (
	"net/http"

	"github.com/go-fuego/fuego"
)
 func GetServicesController(c fuego.ContextNoBody) (*[]Service, error) {
	if(c.Request().Method != http.MethodGet){
		return nil, fuego.BadRequestError{
			Title:  "Invalid ID",
			Detail: "The provided ID is not a valid integer.",
			Err:    nil,
		}
	}
 	return &[]Service{
		{Name: "service 1"},
		{Name: "service 2"},
	}, nil
 }
