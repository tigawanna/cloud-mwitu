package routes

import (
	"github.com/go-fuego/fuego"
)

func RegisterSystemRoutes(s *fuego.Server) {
	group := fuego.Group(s,"/sytemd");
	fuego.Get(group,"/",GetSystemDController)
	fuego.Get(group,"/running",GetRunningSystemDController)
	fuego.Post(group,"/",MakeSystemDController)
}


