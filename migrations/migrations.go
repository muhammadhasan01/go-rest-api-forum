package migrations

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
)

// Migrate is used to migrate the type struct
// to the database
func Migrate() {
	Auth := &interfaces.Auth{}
	User := &interfaces.User{}
	Form := &interfaces.Form{}
	Post := &interfaces.Post{}
	db := utils.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&Auth, &User, &Form, &Post)
}
