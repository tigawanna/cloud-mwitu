package projects

import (
	"net/http"

	"github.com/go-fuego/fuego"
)
 func GetProjectsController(c fuego.ContextNoBody) (*[]Project, error) {
	if(c.Request().Method != http.MethodGet){
		return nil, fuego.BadRequestError{
			Title:  "Invalid ID",
			Detail: "The provided ID is not a valid integer.",
			Err:    nil,
		}
	}
 	return &[]Project{
		{Name: "service 1"},
		{Name: "service 2"},
	}, nil
 }
