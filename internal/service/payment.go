package service

// PaymentServicer 支付服务
type PaymentServicer interface {
}

type PaymentService struct {
}

func NewPaymentService() PaymentServicer {
	s := new(PaymentService)
	return s
}
