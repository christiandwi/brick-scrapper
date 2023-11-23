package server

import (
	"github.com/christiandwi/edot/product-service/interfaces/container"
)

type handler struct {
	productHandler *productHandler
}

func setupHandler(container container.Container) *handler {
	productHandler := newProductHandler(container.ProductService)
	return &handler{
		productHandler: productHandler,
	}
}
