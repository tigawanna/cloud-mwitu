package hello

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {

helloGroup := fuego.Group(s,"/hello");
fuego.Get(helloGroup,"/",HelloGetController)
fuego.Post(helloGroup,"/",HelloPostController)
}





