package thread

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

func GetThreadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response, err := GetThread(uint(key))

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "Thread ID not found"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func AddThreadHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.Thread
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	formattedBody.UserID = claims["user_id"].(uint)
	formattedBody.Username = claims["username"].(string)

	response := AddThread(&formattedBody)

	json.NewEncoder(w).Encode(response)
}

func UpdateThreadHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.Thread
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := UpdateThread(uint(key), formattedBody.Description, claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}

func DeleteThreadHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)

	if err != nil {
		utils.HandleErr(err)
		msg := interfaces.ErrorMessage{ErrorMsg: "ID cannot be converted into an integer"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	response := DeleteThread(uint(key), claims["user_id"].(uint))

	json.NewEncoder(w).Encode(response)
}
