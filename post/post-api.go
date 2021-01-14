package post

import (
	"backend-forum/auth"
	"backend-forum/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Title Gets a post.
// @Description Gets a post from a post ID.
// @Param  threadID  path  int  true  "Thread ID from the path"
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

// @Title Adds a post.
// @Description Adds a post from a post ID.
// @Param  post  body  PostBody  true  "Info of the post (title, description)"
// @Param  threadID  path  int  true  "Thread ID from the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  AddPostResponse  "AddPostResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource post
// @Route /thread/{threadID}/post/add [post]
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Read data from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}

	// Format the data from the body
	var formattedBody PostBody
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		handleError(w, err)
		return
	}

	// Get the threadID from the path
	vars := mux.Vars(r)
	varThreadID, err := strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}
	claims, _ := auth.GetClaims(r)

	threadID := uint(varThreadID)
	// Gets the claims
	userID := claims["user_id"].(uint)
	username := claims["username"].(string)

	response := AddPost(threadID, userID, username, formattedBody.Title, formattedBody.Description)

	json.NewEncoder(w).Encode(response)
}

// @Title Updates a post.
// @Description Updates a post from a post ID.
// @Param  post  body  PostBody  true  "Info of the post (title, description)"
// @Param  threadID  path  int  true  "Thread ID from the path"
// @Param  postID  path  int  true  "Post ID from the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  UpdatePostResponse  "UpdatePostResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource post
// @Route /thread/{threadID}/post/{postID} [put]
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Gets data from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	claims, _ := auth.GetClaims(r)

	// Format the data from the body
	var formattedBody PostBody
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the threadID from the path
	vars := mux.Vars(r)
	_, err = strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the postID from the path
	postID, err := strconv.ParseUint(vars["postID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the response
	response, err := UpdatePost(uint(postID), formattedBody.Title, formattedBody.Description, claims["username"].(string))
	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// @Title Deletes a post.
// @Description Deletes a post from a post ID.
// @Param  threadID  path  int  true  "Thread ID from the path"
// @Param  postID  path  int  true  "Post ID from the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  DeletePostResponse  "DeletePostResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource post
// @Route /thread/{threadID}/post/{postID} [delete]
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Gets claim
	claims, _ := auth.GetClaims(r)
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
	response, err := DeletePost(uint(postID), claims["username"].(string))
	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	utils.HandleErr(err)
	msg := ErrorResponse{Msg: err.Error()}
	json.NewEncoder(w).Encode(msg)
}
