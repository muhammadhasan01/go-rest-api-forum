package migrations

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	"github.com/bxcodec/faker/v3"
)

// CreateAccounts is used to create fake accounts
func CreateAccounts() {
	db := utils.ConnectDB()
	defer db.Close()

	users := []interfaces.User{}
	for i := 0; i < 100; i++ {
		username := faker.Username()
		email := faker.Email()
		password := faker.Password()
		role := "USER"
		user := interfaces.User{
			Username: username,
			Email:    email,
			Password: password,
			Role:     role,
		}
		users = append(users, user)
	}

	for _, user := range users {
		user.Password = utils.HashPassword(user.Password)
		db.Create(&user)
	}
}
