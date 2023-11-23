package repo

import (
	"log"

	"github.com/christiandwi/edot/product-service/database"
	"github.com/christiandwi/edot/product-service/domain/products"
	"github.com/christiandwi/edot/product-service/entity"
)

type productRepo struct {
	db *database.Database
}

func SetupProductRepo(db *database.Database) products.ProductsRepository {
	return &productRepo{db: db}
}

func (p productRepo) InsertProducts(products entity.Products) (err error) {
	if err = p.db.Create(&products).Error; err != nil {
		log.Print("error on inserting products")
	}

	return
}
