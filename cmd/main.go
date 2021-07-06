package main

import (
	"github.com/Aprendendo-programacao/full-cycle-webserver-golang/http"
	"github.com/Aprendendo-programacao/full-cycle-webserver-golang/model"
)

var Products = model.Products{}

func main() {
	webserver := http.NewWebServer()

	webserver.Products = &Products
	webserver.Serve()
}
