package controller

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

// default pagination options
var optionPagination = option.Group(
	option.ResponseHeader("Content-Range", "Total number of caddy configs", 
	param.StatusCodes(200), 
	param.Example("42 caddy configs", "0-10/42")),
)

type CaddyFileResources struct {
	CaddyFileService services.CaddyFileService
}

type CaddyFileError struct {
	Err     error  `json:"-" xml:"-"`
	Message string `json:"message" xml:"message"`
}

var _ error = CaddyFileError{}

func (e CaddyFileError) Error() string { return e.Err.Error() }

func (rs CaddyFileResources) Routes(s *fuego.Server) {
	caddyGroup := fuego.Group(s, "/caddy")

	fuego.Get(caddyGroup, "/", rs.getCaddyFileServices,
		optionPagination,
		option.Query("name", "Filter by name", param.Example("caddy domain name", "localhost"), param.Nullable()),
		option.Description("List all caddyfile services and filter by name"),
	)

	fuego.Get(caddyGroup, "/{name}", rs.getCaddyFileServiceByName,
		option.Description("Get caddyfile service by name"),
	)

	fuego.Post(caddyGroup, "/", rs.updateCaddy,
		option.DefaultStatusCode(201),
		option.Description("Caddyfile will be updated with matching domain record or a new one will be created"),
	)

}

func (rs CaddyFileResources) getCaddyFileServices(c fuego.ContextNoBody) ([]services.CaddyFileModel, error) {
	caddies, err := rs.CaddyFileService.GetCaddyFileItems(c.QueryParam("name"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}

func (rs CaddyFileResources) getCaddyFileServiceByName(c fuego.ContextNoBody) (services.CaddyFileModel, error) {
	caddies, err := rs.CaddyFileService.GetCaddyFileItemByName(c.PathParam("name"))
	if err != nil {
		return services.CaddyFileModel{}, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}

type RequestUpdateCaddyModel struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (rs CaddyFileResources) updateCaddy(c fuego.ContextWithBody[RequestUpdateCaddyModel]) (services.UpdateCaddyResponse, error) {
	body, err := c.Body()
	updateCaddy := services.UpdateCaddyResponse{}
	if err != nil {
		return updateCaddy, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}
	if body.Name == "" || body.Content == "" {
		return updateCaddy, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The body name and content are required",
			Err:    err,
		}
	}
	caddies, err := rs.CaddyFileService.UpdateCaddyFile(body.Name, body.Content)
	if err != nil {
		return updateCaddy, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	err = services.SaveFile("/etc/caddy/Caddyfile", body.Content)
	if err != nil {

		return updateCaddy, fuego.BadRequestError{
			Detail: "Issue saving caddy services file",
			Err:    err,
		}
	}
	return caddies, nil
}
