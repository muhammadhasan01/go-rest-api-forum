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

// @Title Gets as a post.
// @Description Gets a post from a post ID.
// @Param  postID  path  int  true  "postID of the post in the path"
// @Success  200  object  PostResponse  "ThreadResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource post
// @Route /thread/{threadID}/post/{postID} [get]
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Gets the threadID
	_, err := strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the postID
	postID, err := strconv.ParseUint(vars["postID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the response
	response, err := GetPost(uint(postID))
	if err != nil {
		handleError(w, err)
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
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := DeletePost(uint(postID), claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	utils.HandleErr(err)
	msg := ErrorResponse{Msg: err.Error()}
	json.NewEncoder(w).Encode(msg)
}
