package migrations

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	"github.com/bxcodec/faker/v3"
)

func CreateAccounts() {
	db := utils.ConnectDB()
	defer db.Close()

	users := []interfaces.User{}
	for i := 0; i < 100; i++ {
		username := faker.Username()
		email := faker.Email()
		password := faker.Password()
		user := interfaces.User{
			Username: username,
			Email:    email,
			Password: password,
		}
		users = append(users, user)
	}

	for _, user := range users {
		user.Password = utils.HashPassword(user.Password)
		db.Create(&user)
	}
}
