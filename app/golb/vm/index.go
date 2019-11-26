package vm

import (
	"github.com/taadis/letgo/app/golb/models"
)

//
type IndexViewModel struct {
	BaseViewModel
	models.User
	Posts []models.Post
}

//
type IndexViewModelOp struct{}

//
func (vm *IndexViewModelOp) GetVM() IndexViewModel {
	u1, err := models.GetUserByUsername("rene")
	if err != nil {
		panic(err)
	}
	posts, _ := models.GetPostsByUserID(u1.ID)
	v := IndexViewModel{
		BaseViewModel{
			Title: "HomePage",
		},
		u1,
		*posts,
	}
	return v
}
