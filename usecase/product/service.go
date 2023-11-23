package product

type Service interface {
	GetProducts() (err error)
}
