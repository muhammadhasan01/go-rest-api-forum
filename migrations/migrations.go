package migrations

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
)

func Migrate() {
	Auth := &interfaces.Auth{}
	User := &interfaces.User{}
	Form := &interfaces.Form{}
	Post := &interfaces.Post{}
	db := utils.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&Auth, &User, &Form, &Post)

	CreateAccounts()
}
