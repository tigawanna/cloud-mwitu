package routes

import "github.com/go-fuego/fuego"

func RegisterCaddyRoutes(s *fuego.Server) {
	group := fuego.Group(s, "/caddy")
	fuego.Get(group,"/",GetCaddyController)
	// fuego.Get(group,"/running",GetRunningCaddyController)
	// fuego.Post(group,"/",MakeCaddyController)
}
