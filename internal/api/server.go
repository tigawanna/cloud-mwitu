package api

import (
	"github.com/go-fuego/fuego"
	// "net/http"
	// "github.com/go-fuego/fuego/option"
	"github.com/tigawanna/cloud-mwitu/internal/configs"
	"github.com/tigawanna/cloud-mwitu/internal/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/middleware"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

// type NoContent struct {
// 	Result string `json:"result"`
// 	Error  string `json:"error"`
// }

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
			middleware.LogMiddlewereAccess,
			middleware.CorsMiddleware,
		),
		fuego.WithRouteOptions(
			// option.AddResponse(http.StatusNoContent, "No Content",fuego.Response{Type: NoContent{}}),
		))
	s := fuego.NewServer(options...)

	//  caddy endpoints
	caddyService := services.NewCaddyFileService("/path/to/caddy/config")
	caddyfileResorce := controller.CaddyFileResources{
		CaddyFileService: caddyService,
	}
	caddyfileResorce.Routes(s)
	// systemd endpoints
	systemDResource := controller.SystemDFileResources{
		SystemDFileService: services.NewSystemDFileService(),
	}
	systemDResource.Routes(s)

	// fuego.Get(s, "/", func(c fuego.ContextNoBody) (NoContent, error) {
	// 	return NoContent{}, nil
	// })

	return s
}
