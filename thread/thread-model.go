package thread

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	"errors"

	log "github.com/sirupsen/logrus"
)

// AddThread is a function to add a thread to the database
func AddThread(userID uint, username string, name string, description string) AddThreadResponse {
	db := utils.ConnectDB()
	defer db.Close()

	// Defines the thread
	thread := interfaces.Thread{
		UserID:      userID,
		Username:    username,
		Name:        name,
		Description: description,
	}

	// Create the thread to the database
	db.Create(&thread)

	// Log the info
	log.WithFields(log.Fields{
		"threadID": thread.ID,
		"username": username,
	}).Info("A user has just created a new thread")

	// Create the response and return it
	response := AddThreadResponse{
		ID:       thread.ID,
		Message:  "thread has been added successfully!",
		Username: username,
		Name:     name,
	}

	return response
}

// GetThread is a function to get a thread from a thread id
func GetThread(threadID uint) (ThreadResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return ThreadResponse{}, err
	}

	response := ThreadResponse{
		ID:          thread.ID,
		Username:    thread.Username,
		Name:        thread.Name,
		Description: thread.Description,
	}

	return response, nil
}

// UpdateThread is a function to update a thread
func UpdateThread(threadID uint, name string, description string, username string) (UpdateThreadResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the thread exists
	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return UpdateThreadResponse{}, errors.New("thread ID not found")
	}

	// Check whether user can change the thread or not
	if thread.Username != username {
		return UpdateThreadResponse{}, errors.New("You cannot change description of other person thread")
	}

	// Update the thread and save it
	thread.Name = name
	thread.Description = description
	db.Save(&thread)

	// Log the info
	log.WithFields(log.Fields{
		"threadID": thread.ID,
		"username": username,
	}).Info("A user has just updated a thread")

	// Create the response and return it
	response := UpdateThreadResponse{
		Message:     "thread has been updated successfully!",
		Username:    username,
		Name:        name,
		Description: description,
	}

	return response, nil
}

// DeleteThread is a function to delete a thread from a given threadID and username
func DeleteThread(threadID uint, username string) (DeleteThreadResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the thread exists
	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return DeleteThreadResponse{}, errors.New("thread ID not found")
	}

	if thread.Username != username {
		return DeleteThreadResponse{}, errors.New("You cannot change other person thread")
	}

	// Deletes the thread also the post inside the thread
	var post interfaces.Post
	db.Where("thread_id = ?", threadID).Find(&post)
	db.Unscoped().Delete(&post)
	db.Unscoped().Delete(&thread)

	// Log the info
	log.WithFields(log.Fields{
		"threadID": thread.ID,
		"username": username,
	}).Info("A user has just deleted a thread")

	// Create the response and return it
	response := DeleteThreadResponse{
		Message:  "thread has been deleted successfully!",
		ID:       thread.ID,
		Username: username,
	}

	return response, nil
}
