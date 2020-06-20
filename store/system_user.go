package store

// import "github.com/jinzhu/gorm"

// SystemUser
type SystemUser struct {
	Id       string `gorm:"primary_key"`
	Name     string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
}

// TableName
func (*SystemUser) TableName() string {
	return "system_user"
}
