package service

// AreaServicer 地域服务
type AreaServicer interface {
}

type AreaService struct {
}

func NewAreaService() AreaServicer {
	s := new(AreaService)
	return s
}
