package systemd

import (
	"fmt"
	// "net/http"

	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

type SystemDServiceSlice = []services.SystemDService

type SystemDServiceSliceResponse struct {
	Total int                       `json:"total"`
	Items []services.SystemDService `json:"items"`
}

type QueryParams struct {
	Name string `query:"name"`
	Dir  string `query:"dir"`
}
type NoBody struct {
}

func GetSystemDController(c fuego.ContextNoBody) (SystemDServiceSliceResponse, error) {
	queryParams := QueryParams{
		Name: c.QueryParams().Get("name"),
		Dir:  c.QueryParams().Get("dir"),
	}
	servicesList := services.GetSystemDServiceFiles(queryParams.Name, queryParams.Dir == "lib")
	return SystemDServiceSliceResponse{
		Total: len(servicesList),
		Items: servicesList,
	}, nil
	// return &servicesList, nil
}

func GetRunningSystemDController(c fuego.ContextNoBody) (SystemDServiceSliceResponse, error) {

	queryParams := QueryParams{
		Name: c.QueryParams().Get("name"),
		Dir:  c.QueryParams().Get("dir"),
	}
	servicesList := services.GetSystemDServiceFiles(queryParams.Name, queryParams.Dir == "lib")
	return SystemDServiceSliceResponse{
		Total: len(servicesList),
		Items: servicesList,
	}, nil
	// return &servicesList, nil
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
