package service

// CategoryServicer 类目服务
type CategoryServicer interface {
}

type CategoryService struct {
}

func NewCategoryService() CategoryServicer {
	s := new(CategoryService)
	return s
}
