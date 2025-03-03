package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

func RegisterCaddyRoutes(s *fuego.Server) {
	group := fuego.Group(s, "/caddy")
	fuego.Get(group,"/",GetCaddyController)
	fuego.Post(group,"/",RequestUpdateCaddyController)
	fuego.Post(group,"/confirm-update",UpdateCaddyController)
	// fuego.Post(group,"/",MakeCaddyController)
}


func GetCaddyController(c fuego.ContextNoBody) ([]services.CaddyService, error) {
	queryParams := QueryParams{
		Name: c.QueryParams().Get("name"),
	}
	caddies,err:= services.ListCaddyServices(queryParams.Name)
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}
type RequestUpdateCaddyModel struct {
	Name string `json:"name"`
	Content string `json:"content"`
}

func RequestUpdateCaddyController(c fuego.ContextWithBody[RequestUpdateCaddyModel]) (services.UpdateCaddyResponse, error) {
	body, err := c.Body()
	updateCaddy := services.UpdateCaddyResponse{}
	if err != nil {
		return updateCaddy, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}
	if(body.Name=="" || body.Content==""){
		return updateCaddy, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The body name and content are required",
			Err:    err,
		}
	}
	caddies,err:= services.UpdateCaddyFile(body.Name,body.Content)
	if err != nil {
		return updateCaddy, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}
type UpdateCaddyModel struct {
	Content string `json:"content"`
}
type UpdateCaddyResponse struct {
	Status string `json:"status"`
}
func UpdateCaddyController(c fuego.ContextWithBody[UpdateCaddyModel]) (UpdateCaddyResponse, error) {
	updateCaddyResponse := UpdateCaddyResponse{}
	body, err := c.Body()
	if err != nil {
		updateCaddyResponse.Status="Error"
		return updateCaddyResponse, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}
	if(body.Content==""){
		updateCaddyResponse.Status="Error"
		return updateCaddyResponse, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The body content field is required",
			Err:    err,
		}
	}
	err = services.SaveFile("/etc/caddy/Caddyfile",body.Content)
	if err != nil {
		updateCaddyResponse.Status="Error"
		return updateCaddyResponse, fuego.BadRequestError{
			Detail: "Issue saving caddy services file",
			Err:    err,
		}
	}
	updateCaddyResponse.Status="Success"
	return updateCaddyResponse, nil
}
