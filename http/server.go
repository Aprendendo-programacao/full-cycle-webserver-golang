package http

import (
	"github.com/Aprendendo-programacao/full-cycle-webserver-golang/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
	Products *model.Products
	Port     string
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()

	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Start(w.Port)
}

func (w WebServer) getAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(ctx echo.Context) error {
	product := model.NewProduct()

	if err := ctx.Bind(product); err != nil {
		return err
	}

	w.Products.Add(product)

	return ctx.JSON(http.StatusCreated, product)
}
