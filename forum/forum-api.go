package forum

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

func GetForumHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["forumID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response, err := GetForum(uint(key))

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Forum ID not found"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func AddForumHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.Forum
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	formattedBody.UserID = claims["user_id"].(uint)
	formattedBody.Username = claims["username"].(string)

	response := AddForum(&formattedBody)

	json.NewEncoder(w).Encode(response)
}

func UpdateForumHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.Forum
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["forumID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := UpdateForum(uint(key), formattedBody.Description, claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}
