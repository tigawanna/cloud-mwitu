package projects

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	serviceGroup := fuego.Group(s,"/projects");
	fuego.Get(serviceGroup,"/",GetProjectsController)
}
