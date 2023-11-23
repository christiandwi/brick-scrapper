package container

import (
	"github.com/christiandwi/edot/product-service/config"
	"github.com/christiandwi/edot/product-service/database"
	"github.com/christiandwi/edot/product-service/infrastructure/repo"
	"github.com/christiandwi/edot/product-service/usecase/product"
)

type Container struct {
	Config         *config.Config
	ProductService product.Service
}

func SetupContainer() (out Container) {
	cfg := config.SetupConfig()

	db := database.SetupDatabase(cfg)

	productRepo := repo.SetupProductRepo(db)

	//setup asynq package
	// asynq := asynq.NewServer(
	// 	asynq.RedisClientOpt{
	// 		Addr: cfg.Redis.Address,
	// 	},
	// 	asynq.Config{
	// 		Concurrency: cfg.Asynq.Concurrency,
	// 	},
	// )

	//setup product
	productService := product.NewProductService(productRepo)

	productService.GetProducts()

	return Container{
		Config:         cfg,
		ProductService: productService,
	}
}
