package main

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes"
	"github.com/tigawanna/cloud-mwitu/internal/configs"
)



func main() {
	envs:= configs.GetEnv()
	s := fuego.NewServer(
		fuego.WithAddr("localhost:"+envs.Port),	
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(
				fuego.OpenAPIConfig{
					PrettyFormatJSON: true,
				},
			),
		))

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})
	routes.RegisterRoutes(s)
	s.Run()
}
