package services

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	serviceGroup := fuego.Group(s,"/services");
	fuego.Get(serviceGroup,"/",GetServicesController)
}
