package systemd

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	group := fuego.Group(s,"/sytemd");
	fuego.Get(group,"/",GetSystemDController)
	fuego.Get(group,"/running",GetRunningSystemDController)
	fuego.Post(group,"/",MakeSystemDController)
}


