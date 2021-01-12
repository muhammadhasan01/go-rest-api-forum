package forum

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	log "github.com/sirupsen/logrus"
)

func AddForum(forum *interfaces.Forum) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	db.Create(&forum)

	response := map[string]interface{}{"message": "forum added succesfully"}
	log.Info("A new form with the name: ", forum.Name, " has been added succesfully")

	return response
}

func GetForum(forum_id uint) (interfaces.Forum, error) {
	db := utils.ConnectDB()
	defer db.Close()

	var forum interfaces.Forum
	if err := db.First(&forum, forum_id).Error; err != nil {
		return forum, err
	}

	return forum, nil
}

func UpdateForum(forum_id uint, description string, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var forum interfaces.Forum
	if err := db.First(&forum, forum_id).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Forum ID not found"}
	}

	if forum.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot change description of other person forum"}
	}

	forum.Description = description
	db.Save(&forum)

	log.Info("Form with the id ", forum.ID, " has been updated")
	return map[string]interface{}{"message": "forum has been updated succesfully"}
}
