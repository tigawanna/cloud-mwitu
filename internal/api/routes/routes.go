package routes

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	RegisterSystemRoutes(s)
}
