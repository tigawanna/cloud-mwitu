package controller

import (
	"fmt"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

type SystemDFileResources struct {
	SystemDFileService services.SystemDFileService
}

type SystemDFileError struct {
	Err     error  `json:"-" xml:"-"`
	Message string `json:"message" xml:"message"`
}

var _ error = SystemDFileError{}

func (e SystemDFileError) Error() string { return e.Err.Error() }

func (rs SystemDFileResources) Routes(s *fuego.Server) {
	systemDGroup := fuego.Group(s, "/systemd")

	fuego.Get(systemDGroup, "/", rs.getSystemDFileServices,
		option.Query("name", "Filter by name", param.Example("systemd domain name", "apache"), param.Nullable()),
		option.Query("libDir", "look under /lib or /etc", param.Example("systemd directory to look under", "/lib"), param.Nullable()),
		option.Description("List all SystemDFile services and filter by name"),
	)
	fuego.Get(systemDGroup, "/running", rs.getRunningSystemDFileServices,
		option.Query("name", "Filter by name", param.Example("systemd domain name", "postgres"), param.Nullable()),
		option.Query("libDir", "look under /lib or /etc", param.Example("systemd directory to look under", "/lib"), param.Nullable()),
		option.Description("List all Running SystemDFile services and filter by name"),
	)

	fuego.Get(systemDGroup, "/{name}", rs.getSystemDFileServiceByName,
		option.Query("libDir", "look under /lib or /etc", param.Example("systemd directory to look under", "/lib"), param.Nullable()),
		option.Description("Get SystemDFile service by name"),
	)

	fuego.Post(systemDGroup, "/", rs.updateSystemD,
		option.DefaultStatusCode(201),
		option.Description("SystemDFile will be updated with matching domain record or a new one will be created"),
	)

}

func (rs SystemDFileResources) getSystemDFileServices(c fuego.ContextNoBody) ([]services.SystemDService, error) {
	caddies, err := rs.SystemDFileService.GetSystemDServiceFiles(c.QueryParam("name"),c.QueryParam("libDir")!="etc")
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up systemd services",
			Err:    err,
		}
	}
	return caddies, nil
}
func (rs SystemDFileResources) getRunningSystemDFileServices(c fuego.ContextNoBody) ([]services.RunningSystemDService, error) {
	caddies, err := rs.SystemDFileService.GetRunningSystemDServices(c.QueryParam("name"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up systemd services",
			Err:    err,
		}
	}
	return caddies, nil
}

func (rs SystemDFileResources) getSystemDFileServiceByName(c fuego.ContextNoBody) (services.SystemDService, error) {
	caddies, err := rs.SystemDFileService.GetSystemDServiceFiles(c.PathParam("name"),c.QueryParam("libDir")!="etc")
	fmt.Println(" total results ", len(caddies))

	// fmt.Println("getSystemDFileServiceByName === ")
	if err != nil {
		return services.SystemDService{}, fuego.BadRequestError{
			Detail: "Issue looking up systemd services",
			Err:    err,
		}
	}
	if len(caddies) == 0 {
		return services.SystemDService{}, fuego.NotFoundError{
			Detail: "Issue looking up systemd services",
			Err:    err,
		}
	}
	return caddies[0], nil
}

type RequestUpdateSystemDModel struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	LibDir bool `json:"libDir"`
}

func (rs SystemDFileResources) updateSystemD(c fuego.ContextWithBody[RequestUpdateSystemDModel]) (string, error) {
	body, err := c.Body()

	if err != nil {
		return "", fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}
	if body.Name == "" {
		return "", fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The body name and content are required",
			Err:    fmt.Errorf("body name is required"),
		}
	}
	if(body.Content == ""){
		return "", fuego.BadRequestError{
			Title:  "Unexpected data",			
			Detail: "The body content is required",	
			Err:    fmt.Errorf("body content is required"),
		}
	}
	caddies, err := rs.SystemDFileService.UpdateSystemDFile(body.Name, body.Content, body.LibDir)
	if err != nil {
		return "", fuego.BadRequestError{
			Detail: "Issue looking up systemd services",
			Err:    err,
		}
	}


	err = services.SaveFile("/etc/systemd/SystemDFile", body.Content)
	if err != nil {
		return "", fuego.BadRequestError{
			Detail: "Issue saving systemd services file",
			Err:    err,
		}
	}
	return caddies, nil
}
