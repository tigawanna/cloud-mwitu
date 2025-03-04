package api

import (
	"net/http"

	"github.com/go-fuego/fuego"
	// controller "github.com/go-fuego/fuego/examples/petstore/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/configs"
	controller "github.com/tigawanna/cloud-mwitu/internal/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/middleware"
	"github.com/tigawanna/cloud-mwitu/internal/services"

	"github.com/go-fuego/fuego/option"
)

type NoContent struct {
	Empty string `json:"-"`
}

func NewApiServer(options ...func(*fuego.Server)) *fuego.Server {
	options = append(options, 
		fuego.WithAddr("localhost:"+configs.GetEnv().Port),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(
				fuego.OpenAPIConfig{
					PrettyFormatJSON: true,
				},
			),
		),
		fuego.WithGlobalMiddlewares(
			middleware.CorsMiddleware,
			middleware.LogMiddlewereAccess,
		),
		fuego.WithRouteOptions(
		option.AddResponse(http.StatusNoContent, "No Content",
		 fuego.Response{Type: NoContent{}}),
	))
	s := fuego.NewServer(options...)

	caddyService := services.NewCaddyFileService("/path/to/caddy/config")  
	caddyfileResorce :=  controller.CaddyFileResources{
		CaddyFileService:caddyService ,
	}
	caddyfileResorce.Routes(s)

	return s
}
