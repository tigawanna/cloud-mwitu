package hello

import (
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/labstack/echo/v4"
)

type Handler struct {}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World! from /hello")
}

type MyInput struct {
    Name string `json:"name"`
}

type MyHello struct {
    Name string `json:"name"`
}

func HelloPostController(c fuego.ContextWithBody[MyInput]) (*MyHello, error) {
    body, err := c.Body()
    if err != nil {
        return nil, err
    }

    return &MyHello{
        Name: body.Name,
    }, nil
}
func HelloGetController(c fuego.ContextNoBody) (*MyHello, error) {
 	return &MyHello{
        Name: "/hello , i am working",
    }, nil
}
