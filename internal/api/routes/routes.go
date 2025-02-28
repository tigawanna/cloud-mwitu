package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes/hello"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes/systemd"
)

func RegisterRoutes(s *fuego.Server) {
    // Register route groups
    hello.RegisterRoutes(s)
	systemd.RegisterRoutes(s)
}
