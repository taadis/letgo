package service

// UserServicer 用户服务
type UserServicer interface {
}

type UserService struct {
}

func NewUserService() UserServicer {
	s := new(UserService)
	return s
}
