package post

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"errors"

	log "github.com/sirupsen/logrus"
)

// AddPost is a functio to add a post
func AddPost(threadID uint, userID uint, username string, title string, description string) (AddPostResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the thread exists
	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return AddPostResponse{}, errors.New("thread ID not found")
	}

	// Make the post
	post := interfaces.Post{
		ThreadID:    threadID,
		UserID:      userID,
		Username:    username,
		Title:       title,
		Description: description,
	}

	// Add the post to the database
	db.Create(&post)

	// Log the info
	log.WithFields(log.Fields{
		"postID":   post.ID,
		"username": username,
	}).Info("A user has just created a new post")

	// Create the response and return it
	response := AddPostResponse{
		ID:       post.ID,
		Message:  "post has been added successfully!",
		Username: username,
		Title:    title,
	}

	return response, nil
}

// GetPost is a function to get a post from a post ID
func GetPost(threadID uint, postID uint) (PostResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the post exists
	var post interfaces.Post
	if err := db.Where("id = ? AND thread_id = ? ", postID, threadID).First(&post, postID).Error; err != nil {
		return PostResponse{}, err
	}

	response := PostResponse{
		ID:          post.ID,
		Username:    post.Username,
		Title:       post.Title,
		Description: post.Description,
	}

	return response, nil
}

// UpdatePost is a function to update a post
func UpdatePost(threadID uint, postID uint, title string, description string, username string) (UpdatePostResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the thread exists
	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return UpdatePostResponse{}, errors.New("thread ID not found")
	}

	// Check whether the post exists
	var post interfaces.Post
	if err := db.First(&post, postID).Error; err != nil {
		return UpdatePostResponse{}, errors.New("post ID not found")
	}

	// Check whether user can change the post or not
	if post.Username != username {
		return UpdatePostResponse{}, errors.New("You cannot change description of other person post")
	}

	// Update the post and save it
	post.Title = title
	post.Description = description
	db.Save(&post)

	// Log the info
	log.WithFields(log.Fields{
		"postID":   post.ID,
		"username": username,
	}).Info("A user has just updated a post")

	// Create the response and return it
	response := UpdatePostResponse{
		Message:     "post has been updated successfully!",
		Username:    username,
		Title:       title,
		Description: description,
	}

	return response, nil
}

// DeletePost is a function to delete a post
func DeletePost(threadID uint, postID uint, username string) (DeletePostResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the thread exists
	var thread interfaces.Thread
	if err := db.First(&thread, threadID).Error; err != nil {
		return DeletePostResponse{}, errors.New("thread ID not found")
	}

	// Check whether the post exists
	var post interfaces.Post
	if err := db.First(&post, postID).Error; err != nil {
		return DeletePostResponse{}, errors.New("post ID not found")
	}

	if post.Username != username {
		return DeletePostResponse{}, errors.New("You cannot delete other person post")
	}

	// Deletes the post
	db.Unscoped().Delete(&post)

	// Log the info
	log.WithFields(log.Fields{
		"postID":   post.ID,
		"username": username,
	}).Info("A user has just deleted a post")

	// Create the response and return it
	response := DeletePostResponse{
		Message:  "post has been deleted successfully!",
		ID:       post.ID,
		Username: username,
	}

	return response, nil
}
