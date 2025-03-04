package controller

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

// default pagination options
var optionPagination = option.Group(
	option.QueryInt("per_page", "Number of items per page", param.Required()),
	option.QueryInt("page", "Page number", param.Default(1), param.Example("1st page", 1), param.Example("42nd page", 42), param.Example("100th page", 100)),
	option.ResponseHeader("Content-Range", "Total number of caddy configs", param.StatusCodes(200, 206), param.Example("42 caddy configs", "0-10/42")),
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
	caddyGroup := fuego.Group(s, "/caddy", option.Header("X-Header", "header description"))
	
	fuego.Get(caddyGroup, "/", rs.getCaddyFileServicesController,
		optionPagination,
		option.Query("name", "Filter by name", param.Example("caddy domain name", "localhost"), param.Nullable()),
		option.QueryInt("younger_than", "Only get pets younger than given age in years", param.Default(3)),
		option.Description("Filter caddy services by name"),
	)

	fuego.Get(caddyGroup, "/all", rs.getCaddyFileServiceByNameController,
		optionPagination,
		option.Query("name", "Name of the service", param.Required(), param.Example("example 1", "Napoleon")),
		option.Tags("my-tag"),
		option.Description("Get all pets"),
	)

	fuego.Post(caddyGroup, "/", rs.updateCaddyController,
		option.DefaultStatusCode(201),
		option.AddResponse(409, "Conflict: Pet with the same name already exists", fuego.Response{Type: CaddyFileError{}}),
	)

}



func (rs CaddyFileResources) getCaddyFileServicesController(c fuego.ContextNoBody) ([]services.CaddyFileModel, error) {
	caddies, err := rs.CaddyFileService.GetCaddyFileItems(c.QueryParam("name"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}

func (rs CaddyFileResources) getCaddyFileServiceByNameController(c fuego.ContextNoBody) (services.CaddyFileModel, error) {
	caddies, err := rs.CaddyFileService.GetCaddyFileItemByName(c.QueryParam("name"))
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



func (rs CaddyFileResources) updateCaddyController(c fuego.ContextWithBody[RequestUpdateCaddyModel]) (services.UpdateCaddyResponse, error) {
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










