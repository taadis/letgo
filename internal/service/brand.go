package service

// BrandServicer 品牌服务
type BrandServicer interface {
}

type BrandService struct {
}

func NewBrandService() BrandServicer {
	s := new(BrandService)
	return s
}
