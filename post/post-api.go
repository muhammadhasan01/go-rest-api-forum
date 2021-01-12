package post

import (
	"backend-forum/auth"
	"backend-forum/interfaces"
	"backend-forum/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, err := strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Thread ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	postID, err := strconv.ParseUint(vars["postID"], 10, 64)
	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Post ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response, err := GetPost(uint(postID))

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Post ID not found"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	var formattedBody interfaces.Post
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	vars := mux.Vars(r)
	threadID, err := strconv.ParseUint(vars["threadID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Thread ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	claims, _ := auth.GetClaims(r)

	formattedBody.ThreadID = uint(threadID)
	formattedBody.UserID = claims["user_id"].(uint)
	formattedBody.Username = claims["username"].(string)

	response := AddPost(&formattedBody)

	json.NewEncoder(w).Encode(response)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.Post
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	vars := mux.Vars(r)
	_, err = strconv.ParseUint(vars["threadID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Thread ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	postID, err := strconv.ParseUint(vars["postID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := UpdatePost(uint(postID), formattedBody.Description, claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["postID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := DeletePost(uint(key), claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}
