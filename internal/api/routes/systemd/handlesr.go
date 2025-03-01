package systemd

import (
	"fmt"
	// "net/http"

	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)


type SystemDServiceSlice = []services.SystemDService

type SystemDServiceSliceResponse struct {
	Total int 	`json:"total"`
	Items []services.SystemDService `json:"items"`
}

func GetSystemDController(c fuego.ContextNoBody) (SystemDServiceSliceResponse, error) {	
	servicesList := services.GetSystemDServiceFiles("",true)
	return SystemDServiceSliceResponse{
		Total:   len(servicesList),
		Items:   servicesList,
	},nil
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
	config := services.NewSystemdServiceConfig("my-service", "~/myapp", "myapp", nil)
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
		ServiceFile: content,
		CreatedOrUpdated: "created",
		}, nil

}
