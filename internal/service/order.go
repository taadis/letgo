package service

// OrderServicer 订单服务
type OrderServicer interface {
}

type OrderService struct {
}

func NewOrderService() OrderServicer {
	s := new(OrderService)
	return s
}
