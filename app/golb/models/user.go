package models

import "fmt"

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePassword(password)
}

func GetUserByUsername(username string) User {
	var user User
	fmt.Println(db)
	db.Where("username=?", username).Find(&user)
	return user
}

func AddUser(username, password, email string) error {
	user := User{
		Username: username,
		Email:    email,
	}
	user.SetPassword(password)
	return db.Create(&user).Error
}
