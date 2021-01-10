package interfaces

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username string
	Email    string
	Password string
}

type Auth struct {
	gorm.Model

	UserID uint
	Token  string
}

type Form struct {
	gorm.Model

	UserID      uint
	Name        string
	Description string
}

type Post struct {
	gorm.Model

	UserID      uint
	FormID      uint
	Title       string
	Description string
}
