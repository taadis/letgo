package common

// Paging
type Paging struct {
	Page  int `json:"page" binding:"-"`
	Limit int `json:"limit" binding:"-"`
}
