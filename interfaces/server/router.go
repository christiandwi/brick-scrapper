package server

import "github.com/gin-gonic/gin"

func setupRouter(app *gin.Engine, handler *handler) {
	// appRouter := app.Group("/api")
	// v1 := appRouter.Group("/v1")

	//Product grouping
	// product := v1.Group("/products")
	// product.GET("", handler.productHandler.GetProducts())
}
