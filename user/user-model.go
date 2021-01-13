package user

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	log "github.com/sirupsen/logrus"
)

func GetUser(username string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var user interfaces.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "username not found"}
	}

	response := map[string]interface{}{
		"userID":   user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	return response
}

func UpdateUser(username string, password string, username_claim string) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var user interfaces.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "User ID not found"}
	}

	if user.Role == "USER" && user.Username != username_claim {
		return map[string]interface{}{"ErrorMsg": "As a Role USER you cannot change other person password"}
	}

	user.Password = utils.HashPassword(password)
	db.Save(&user)

	log.Info("User with the username ", user.Username, " has been updated")
	return map[string]interface{}{"message": "user has been updated succesfully"}
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
