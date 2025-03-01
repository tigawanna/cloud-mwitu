package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

func GetCaddyController(c fuego.ContextNoBody) ([]services.CaddyService, error) {
	caddies,err:= services.ListCaddyServices("")
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Issue looking up caddy services",
			Err:    err,
		}
	}
	return caddies, nil
}
