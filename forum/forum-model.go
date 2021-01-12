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
