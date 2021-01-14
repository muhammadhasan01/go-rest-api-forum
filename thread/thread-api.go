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

// @Title Gets as a thread.
// @Description Gets a thread from a thread ID.
// @Param  threadID  path  int  true  "threadID of the thread in the path"
// @Success  200  object  ThreadResponse  "ThreadResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource thread
// @Route /thread/{threadID} [get]
func GetThreadHandler(w http.ResponseWriter, r *http.Request) {
	// Gets the threadID from the path
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)

	// Check if there is error in the threadID
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets response
	response, err := GetThread(uint(key))

	// Check if there are any bad request happenned
	if err != nil {
		handleError(w, err)
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

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	utils.HandleErr(err)
	msg := ErrorResponse{Msg: err.Error()}
	json.NewEncoder(w).Encode(msg)
}
