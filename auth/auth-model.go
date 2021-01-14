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

	// Logs out the info and return the response
	log.WithFields(log.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("A user has just logged in")

	response := LoginResponse{Message: "you have been logged in succesfully!", Token: token}

	return response, nil
}

// Register function is used handle
// register a user, it needs username
// email and password to store it in the database
func Register(username string, email string, pass string) (RegisterResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Generate the password with hashing
	generatedPassword := utils.HashPassword(pass)
	user := &interfaces.User{Username: username, Email: email, Password: generatedPassword, Role: "USER"}

	// Check if the username has already been taken
	if !db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return RegisterResponse{}, errors.New("username has already been taken")
	}

	// Check if the email has already been taken
	if !db.Where("email = ? ", email).First(&user).RecordNotFound() {
		return RegisterResponse{}, errors.New("email has already been taken")
	}

	// Create a new record in the database
	db.Create(&user)

	// Logs out the info and return the response
	log.WithFields(log.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("A user has just registered")

	response := RegisterResponse{
		Message:  "user registered successfully!",
		Username: user.Username,
		Email:    user.Email,
	}

	return response, nil
}

// Logout function is to handle when a user logout
// it deletes the record database in the whitelist
func Logout(user_id uint, token string, username string) LogoutResponse {
	db := utils.ConnectDB()
	defer db.Close()

	// Deletes the record in the database whitelist
	auth := &interfaces.Auth{}
	db.Where(map[string]interface{}{"user_id": user_id, "token": token}).First(&auth)
	db.Unscoped().Delete(&auth)

	// Logs out the info and return the response
	log.WithFields(log.Fields{
		"user_id":  user_id,
		"username": username,
	}).Info("A user has just logged out")

	response := LogoutResponse{
		Message:  "you have been logout successfully!",
		Username: username,
	}

	return response
}
