package main

import (
	"github.com/christiandwi/edot/product-service/interfaces/container"
	"github.com/christiandwi/edot/product-service/interfaces/server"
)

func main() {
	server.Start(container.SetupContainer())
}
