package products

import "github.com/christiandwi/edot/product-service/entity"

type ProductsRepository interface {
	InsertProducts(products entity.Products) (err error)
}
