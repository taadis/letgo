package service

type ProductServicer interface {
}

// ProductService 商品服务
type ProductService struct {
}

func NewProductService() ProductServicer {
	s := new(ProductService)
	return s
}
