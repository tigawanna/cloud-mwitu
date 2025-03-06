package api

import (
	"github.com/go-fuego/fuego"
	// "net/http"
	// "github.com/go-fuego/fuego/option"
	"github.com/tigawanna/cloud-mwitu/internal/configs"
	"github.com/tigawanna/cloud-mwitu/internal/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/db"
	"github.com/tigawanna/cloud-mwitu/internal/middleware"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

// type NoContent struct {
// 	Result string `json:"result"`
// 	Error  string `json:"error"`
// }

func NewApiServer(options ...func(*fuego.Server)) *fuego.Server {

  // Initialize database
    database := db.InitDB()
    
    // Initialize services that require database access
    authService := services.NewAuthService(database, nil)
    
	
	
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
			middleware.AuthMiddleware(authService),
			middleware.CorsMiddleware,
			middleware.LogMiddlewereAccess,
			// Create the auth middleware with the auth service
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

	// Use default session config
	authController := controller.AuthResources{
    AuthService: authService,
	}

	// Register auth routes
	authController.Routes(s)

	// fuego.Get(s, "/", func(c fuego.ContextNoBody) (NoContent, error) {
	// 	return NoContent{}, nil
	// })

	return s
}
