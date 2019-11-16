package vm

import (
	"github.com/taadis/letgo/app/golb/models"
)

//
type IndexViewModel struct {
	BaseViewModel
	Models.User
	Posts []models.Post
}

//
type IndexViewModelOp struct{}

//
func (vm *IndexViewModelOp) GetVM() IndexViewModel {
	u1 := models.GetUserByUsername("rene")
	posts, _ := models.GetPostsByUserID(u1.ID
	v := IndexViewModel{
		BaseViewModel{
			Title:"HomePage",
		},
		u1,
		*posts,
	}
	return v
}
