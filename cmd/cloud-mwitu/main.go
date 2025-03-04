package main

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes"
	"github.com/tigawanna/cloud-mwitu/internal/configs"
	"github.com/tigawanna/cloud-mwitu/internal/middleware"
)

func main() {
	envs := configs.GetEnv()
	s := fuego.NewServer(
		fuego.WithAddr("localhost:"+envs.Port),
		fuego.WithGlobalMiddlewares(
			middleware.CorsMiddleware,
			middleware.LogMiddlewereAccess,
		),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(
				fuego.OpenAPIConfig{
					PrettyFormatJSON: true,
				},
			),
		))
	type Welcome struct {
		Message string `json:"message"`
	}
	fuego.Get(s, "/", func(c fuego.ContextNoBody) (Welcome, error) {
		return Welcome{Message: "Hello World"}, nil
	})
	fuego.Get(s, "/caddre", routes.GetCaddyController)
	routes.RegisterRoutes(s)
	s.Run()
}
