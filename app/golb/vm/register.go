// register.go
package vm

import (
	"github.com/taadis/letgo/app/golb/models"
)

type RegisterViewModel struct {
	BaseViewModel
}

type RegisterViewModelOp struct {
}

func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserName(username string) bool {
	_, err := models.GetUserByUsername(username)
	if err != nil {
		return false
	}
	return true
}

func AddUser(username, password, email string) error {
	return models.AddUser(username, password, email)
}
