package post

import (
	"backend-forum/interfaces"
	"backend-forum/utils"

	log "github.com/sirupsen/logrus"
)

func AddPost(Post *interfaces.Post) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	db.Create(&Post)

	response := map[string]interface{}{"message": "Post added succesfully", "post": Post}
	log.Info("A new post with the title: ", Post.Title, " has been added succesfully")

	return response
}

// GetPost is a function to get a post from a post ID
func GetPost(postID uint) (PostResponse, error) {
	db := utils.ConnectDB()
	defer db.Close()

	// Check whether the post exists
	var post interfaces.Post
	if err := db.First(&post, postID).Error; err != nil {
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

func UpdatePost(postID uint, description string, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var Post interfaces.Post
	if err := db.First(&Post, postID).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Post ID not found"}
	}

	if Post.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot change description of other person Post"}
	}

	Post.Description = description
	db.Save(&Post)

	log.Info("Post with the id ", Post.ID, " has been updated")
	return map[string]interface{}{"message": "Post has been updated succesfully", "newPost": Post}
}

func DeletePost(postID uint, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var Post interfaces.Post
	if err := db.First(&Post, postID).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Post ID not found"}
	}

	if Post.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot delete other person Post"}
	}

	db.Unscoped().Delete(&Post)

	log.Info("Post with the id ", Post.ID, " has been deleted")
	return map[string]interface{}{"message": "Post has been deleted succesfully"}
}
