package server

import (
	"github.com/christiandwi/edot/product-service/usecase/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func newProductHandler(service product.Service) *productHandler {
	return &productHandler{service: service}
}

func (g *productHandler) GetProducts() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if err := g.service.GetProducts(); err != nil {
			panic(err)
		}
	}
}
