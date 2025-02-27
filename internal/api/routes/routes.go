package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes/hello"
)

func RegisterRoutes(s *fuego.Server) {
    // Register route groups
    hello.RegisterRoutes(s)
}
