package interfaces

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username string
	Email    string
	Password string
	Role     string
}

type Auth struct {
	gorm.Model

	UserID uint
	Token  string
}

type Thread struct {
	gorm.Model

	UserID      uint
	Username    string
	Name        string
	Description string
}

type Post struct {
	gorm.Model

	UserID      uint
	Username    string
	ThreadID    uint
	Title       string
	Description string
}

type ErrorMessage struct {
	ErrorMsg string
}
