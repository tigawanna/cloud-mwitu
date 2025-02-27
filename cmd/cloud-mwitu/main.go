package main

import (
	"github.com/go-fuego/fuego"
	"github.com/tigawanna/cloud-mwitu/internal/api/routes"
)


type MyInput struct {
	Name string `json:"name" validate:"required"`
}

type MyOutput struct {
	Message string `json:"message"`
}


// func myController(c fuego.ContextWithBody[MyInput]) (*MyOutput, error) {
// 	body, err := c.Body()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MyOutput{Message: "Hello, " + body.Name}, nil
// }


func main() {
	s := fuego.NewServer(
		fuego.WithAddr("localhost:8080"),
		
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
	// fuego.Post(s, "/user/{user}", myController)
	routes.RegisterRoutes(s)

	s.Run()
}
