package routes

import (
	"fmt"
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)


type QueryParams struct {
	Name string `query:"name"`
	Dir  string `query:"dir"`
}

type SystemDServiceSlice = []services.SystemDService

type SystemDServiceResponse struct {
	Total int                 `json:"total"`
	Items SystemDServiceSlice `json:"items"`
}

func GetSystemDController(c fuego.ContextNoBody) (SystemDServiceResponse, error) {
	queryParams := QueryParams{
		Name: c.QueryParams().Get("name"),
		Dir:  c.QueryParams().Get("dir"),
	}
	servicesList := services.GetSystemDServiceFiles(queryParams.Name, queryParams.Dir == "lib")
	return SystemDServiceResponse{
		Total: len(servicesList),
		Items: servicesList,
	}, nil
	// return &servicesList, nil
}

type RunningSystemDServiceSlice = []services.RunningSystemDService

type RunningSystemDServiceResponse struct {
	Total int                        `json:"total"`
	Items RunningSystemDServiceSlice `json:"items"`
}

func GetRunningSystemDController(c fuego.ContextNoBody) (RunningSystemDServiceResponse, error) {

	queryParams := QueryParams{
		Name: c.QueryParams().Get("name"),
	}
	servicesList := services.GetRunningSystemDServices(queryParams.Name)
	return RunningSystemDServiceResponse{
		Total: len(servicesList),
		Items: servicesList,
	}, nil
	// return &servicesList, nil
}




type CreateSystemDModel struct {
	ServiceName string `json:"serviceName"`
	BaseDir  string `json:"baseDir"`
	ExecCommand string `json:"execCommand"`
	LibDir	bool `json:"libDir"`
}
type CreateSystemDResponseModel struct {
	CreateSystemDModel `json:"createSystemDModel"`
	ServiceFile string `json:"serviceFile"`
	CreatedOrUpdated string `json:"createdOrUpdated"`
}


func MakeSystemDController(c fuego.ContextWithBody[CreateSystemDModel]) (*CreateSystemDResponseModel, error) {
	body, err := c.Body()
	if err != nil {
		return nil, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}
	config := services.NewSystemdServiceConfig(body.ServiceName, body.BaseDir, body.ExecCommand, body.LibDir, nil)
	content, err := config.ToString()
	if err != nil {
		fmt.Println("Error generating service file:", err)
		return nil, fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "rror generating service file:",
			Err:    err,
		}
	}
	return &CreateSystemDResponseModel{
		CreateSystemDModel: body,
		ServiceFile:        content,
		CreatedOrUpdated:   "created",
	}, nil

}
