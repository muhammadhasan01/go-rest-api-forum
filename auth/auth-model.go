package auth

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"errors"

	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

// Login handles the user to login
// it checks the username and its password
// also gives out token
func Login(username string, pass string) (LoginResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check the username
	user := &interfaces.User{}
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return LoginResponse{}, errors.New("user not found")
	}

	// Check the password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return LoginResponse{}, errors.New("password is wrong")
	}

	// Check the token first
	token := GetToken(user)

	if CheckTokenInDB(token) {
		return LoginResponse{}, errors.New("User has already logged in")
	}

	// Creates a new record in the "whitelist"
	auth := &interfaces.Auth{UserID: user.ID, Token: token}
	db.Create(&auth)

	log.WithFields(log.Fields{
		"username": username,
	}).Info("A user has just logged in")
	// log.Info("User with the username:", username, " has just logged in")
	return LoginResponse{Message: "you have been logged in succesfully!", Token: token}, nil
}

func Register(username string, email string, pass string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	generatedPassword := utils.HashPassword(pass)
	user := &interfaces.User{Username: username, Email: email, Password: generatedPassword, Role: "USER"}

	if !db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"success": false, "message": "username has already been taken"}
	}

	if !db.Where("email = ? ", email).First(&user).RecordNotFound() {
		return map[string]interface{}{"success": false, "message": "email has already been taken"}
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
