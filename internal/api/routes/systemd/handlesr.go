package systemd

import (
	"fmt"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services/systemd"
)

func GetSystemDController(c fuego.ContextNoBody) (*[]CreateSystemDModel, error) {
	if c.Request().Method != http.MethodGet {
		return nil, fuego.BadRequestError{
			Title:  "Invalid ID",
			Detail: "The provided ID is not a valid integer.",
			Err:    nil,
		}
	}
	return &[]CreateSystemDModel{
		{ServiceName: "service 1", BaseDir: "test", ExecCommand: "test"},
		{ServiceName: "service 2", BaseDir: "test", ExecCommand: "test"},
	}, nil
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
	config := systemd.NewSystemdServiceConfig("my-service", "~/myapp", "myapp", nil)
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
