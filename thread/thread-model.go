package thread

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	log "github.com/sirupsen/logrus"
)

func AddThread(thread *interfaces.Thread) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	db.Create(&thread)

	response := map[string]interface{}{"message": "thread added succesfully"}
	log.Info("A new form with the name: ", thread.Name, " has been added succesfully")

	return response
}

func GetThread(thread_id uint) (interfaces.Thread, error) {
	db := utils.ConnectDB()
	defer db.Close()

	var thread interfaces.Thread
	if err := db.First(&thread, thread_id).Error; err != nil {
		return thread, err
	}

	return thread, nil
}

func UpdateThread(thread_id uint, description string, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var thread interfaces.Thread
	if err := db.First(&thread, thread_id).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Thread ID not found"}
	}

	if thread.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot change description of other person thread"}
	}

	thread.Description = description
	db.Save(&thread)

	log.Info("Thread with the id ", thread.ID, " has been updated")
	return map[string]interface{}{"message": "thread has been updated succesfully"}
}

func DeleteThread(thread_id uint, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var thread interfaces.Thread
	if err := db.First(&thread, thread_id).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Thread ID not found"}
	}

	if thread.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot delete other person thread"}
	}

	db.Unscoped().Delete(&thread)

	log.Info("Thread with the id ", thread.ID, " has been deleted")
	return map[string]interface{}{"message": "thread has been deleted succesfully"}
}