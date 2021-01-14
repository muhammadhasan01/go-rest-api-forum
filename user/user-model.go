package user

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"errors"

	log "github.com/sirupsen/logrus"
)

// GetUser is a function to get a user from a username
func GetUser(username string) (UserResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the username exist in the database
	var user interfaces.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return UserResponse{}, errors.New("username not found")
	}

	// Return the response
	response := UserResponse{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return response, nil
}

// UpdateUser is a function to update a password user from a username
func UpdateUser(username string, password string, username_claim string) (UpdateResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check if username exists
	var user interfaces.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return UpdateResponse{}, errors.New("username not found")
	}

	// Find the user claim in the database
	var userClaim interfaces.User
	db.Where("username = ?", username_claim).First(&userClaim)

	// Check whether this user can change the password or not
	if userClaim.Role == "USER" && user.Username != username_claim {
		return UpdateResponse{}, errors.New("as a ROLE user you cannot change other password")
	}

	// Save the new password to the database
	user.Password = utils.HashPassword(password)
	db.Save(&user)

	// Logs out the info and return the response
	log.WithFields(log.Fields{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("A user has just updated its password")

	response := UpdateResponse{
		Message:  "user has been updated with the new password",
		Username: username,
	}

	return response, nil
}

func DeleteUser(username string, username_claim string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var user interfaces.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "User ID not found"}
	}

	if user.Role == "USER" && user.Username != username_claim {
		return map[string]interface{}{"ErrorMsg": "As a Role USER you cannot delete other person user"}
	}

	auth := &interfaces.Auth{}
	db.Where("user_id = ?", user.ID).First(&auth)
	db.Unscoped().Delete(&auth)
	db.Unscoped().Delete(&user)

	log.Info("User with the username ", user.Username, " has been deleted")
	return map[string]interface{}{"message": "user has been deleted succesfully"}
}
