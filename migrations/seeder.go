package migrations

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"math/rand"

	faker "github.com/bxcodec/faker/v3"
)

// CreateAccounts is used to create fake accounts
func CreateAccounts() {
	db := utils.ConnectDB()
	defer db.Close()

	users := []interfaces.User{}

	users = append(users, interfaces.User{
		Username: "hasan",
		Email:    "hasan@gmail.com",
		Password: "password",
		Role:     "SUPERUSER",
	})
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

// CreateThreads is used to create fake thread
func CreateThreads() {
	db := utils.ConnectDB()
	defer db.Close()
	for i := 0; i < 100; i++ {
		username := "hasan"
		userID := uint(1)
		name := faker.Word()
		description := faker.Sentence()
		thread := interfaces.Thread{
			Username:    username,
			UserID:      userID,
			Name:        name,
			Description: description,
		}
		db.Create(&thread)
	}
}

// CreatePosts is used to create fake thread
func CreatePosts() {
	db := utils.ConnectDB()
	defer db.Close()
	for i := 0; i < 100; i++ {
		min := 1
		max := 100
		username := "hasan"
		userID := uint(1)
		threadID := uint(rand.Intn(max-min) + min)
		title := faker.Word()
		description := faker.Sentence()
		post := interfaces.Post{
			Username:    username,
			UserID:      userID,
			ThreadID:    threadID,
			Title:       title,
			Description: description,
		}
		db.Create(&post)
	}
}
