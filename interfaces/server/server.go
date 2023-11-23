package server

import (
	"github.com/christiandwi/edot/product-service/interfaces/container"
)

func Start(container container.Container) {
	// addr := flag.String("addr: ", container.Config.App.Addr, "Address to listen and serve")
	// app := gin.Default()

	// Setup Handler
	setupHandler(container)

	// Setup Router
	// setupRouter(app, handler)

	// fmt.Println(app.Run(*addr))
}
