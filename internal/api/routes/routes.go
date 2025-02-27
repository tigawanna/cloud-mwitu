package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/tigawanna/cloud-mwitu/internal/api/routes/hello"
)

func RegisterRoutes(e *echo.Echo) {
    // Register route groups
    hello.RegisterRoutes(e)
}
