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

	response := map[string]interface{}{"message": "Post added succesfully"}
	log.Info("A new post with the title: ", Post.Title, " has been added succesfully")

	return response
}

func GetPost(Post_id uint) (interfaces.Post, error) {
	db := utils.ConnectDB()
	defer db.Close()

	var Post interfaces.Post
	if err := db.First(&Post, Post_id).Error; err != nil {
		return Post, err
	}

	return Post, nil
}

func UpdatePost(post_id uint, description string, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var Post interfaces.Post
	if err := db.First(&Post, post_id).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Post ID not found"}
	}

	if Post.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot change description of other person Post"}
	}

	Post.Description = description
	db.Save(&Post)

	log.Info("Post with the id ", Post.ID, " has been updated")
	return map[string]interface{}{"message": "Post has been updated succesfully"}
}

func DeletePost(Post_id uint, user_id uint) map[string]interface{} {
	db := utils.ConnectDB()
	defer db.Close()

	var Post interfaces.Post
	if err := db.First(&Post, Post_id).Error; err != nil {
		return map[string]interface{}{"ErrorMsg": "Post ID not found"}
	}

	if Post.UserID != user_id {
		return map[string]interface{}{"ErrorMsg": "You cannot delete other person Post"}
	}

	db.Unscoped().Delete(&Post)

	log.Info("Post with the id ", Post.ID, " has been deleted")
	return map[string]interface{}{"message": "Post has been deleted succesfully"}
}
