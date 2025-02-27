package hello

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	h := NewHandler()
	e.GET("/", h.Hello)
}
