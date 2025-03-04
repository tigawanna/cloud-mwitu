package main

import (
	"github.com/tigawanna/cloud-mwitu/internal/api"
)

func main() {

	s := api.NewApiServer()

	s.Run()
}
