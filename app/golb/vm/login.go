package vm

type BaseViewModel struct {
	Title string
}

func (m *BaseViewModel) SetTitle(value string) {
	m.Title = value
}

type LoginViewModel struct {
	BaseViewModel
}

// LoginViewModelOp strutc
type LoginViewModelOp struct{}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
