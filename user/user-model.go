package user

import (
	"backend-forum/auth"
	"backend-forum/interfaces"
	"backend-forum/utils"

	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	user := &interfaces.User{}
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}

	token := auth.GetToken(user)

	log.Info("User with the username:", username, " has just logged in")
	return map[string]interface{}{"message": "you have been logged in succesfully!", "token": token}
}

func Register(username string, email string, pass string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	generatedPassword := utils.HashPassword(pass)
	user := &interfaces.User{Username: username, Email: email, Password: generatedPassword, Role: "USER"}

	if !db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "username has already been taken"}
	}

	db.Create(&user)

	response := map[string]interface{}{"success": true, "message": "user registered successfully"}

	return response
}

func Logout(user_id uint, token string, username string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	auth := &interfaces.Auth{}
	db.Where(map[string]interface{}{"user_id": user_id, "token": token}).First(&auth)
	db.Unscoped().Delete(&auth)

	log.Info("User with the username:", username, " has just logged out")
	return map[string]interface{}{"message": "you have been logout successfully!"}
}
