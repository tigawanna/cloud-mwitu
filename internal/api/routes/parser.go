package routes

import (
	"github.com/go-fuego/fuego"
)

type TextToParse struct {
	Text string `json:"text"`
}

func RegisterParserRoutes(s *fuego.Server) {
	group := fuego.Group(s, "/parser")
	fuego.Post(group, "/", ParseStringController)
}

type KDlString struct {
	Text string `json:"text"`
}
func ParseStringController(c fuego.ContextWithBody[TextToParse]) (string, error) {

	_, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Title:  "Unexpected data",
			Detail: "The data recieved was unexpected",
			Err:    err,
		}
	}



	return  "parsed", nil

}

