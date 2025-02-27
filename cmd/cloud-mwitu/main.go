package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes"
)

func main(){
	e := echo.New()
	routes.RegisterRoutes(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
