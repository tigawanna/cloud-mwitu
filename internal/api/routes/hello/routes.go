package hello

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {

helloGroup := fuego.Group(s,"/hello");
fuego.Get(helloGroup,"/",HelloGetController)
fuego.Post(helloGroup,"/",HelloPostController)
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
