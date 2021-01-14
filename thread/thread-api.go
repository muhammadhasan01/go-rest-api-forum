package thread

import (
	"backend-forum/auth"
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

// @Title Adds as a thread.
// @Description Adds a thread from a thread ID.
// @Param  thread  body  ThreadBody  true  "Info of the thread (name, description)"
// @Param  threadID  path  int  true  "threadID of the thread in the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  AddThreadResponse  "AddThreadResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource thread
// @Route /thread/{threadID} [post]
func AddThreadHandler(w http.ResponseWriter, r *http.Request) {
	// Gets the data from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the claim
	claims, _ := auth.GetClaims(r)

	// Format the data from the body
	var formattedBody ThreadBody
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		handleError(w, err)
		return
	}

	// Initilaize the claim into variables
	userID := claims["user_id"].(uint)
	username := claims["username"].(string)

	// Get response and give out the response
	response := AddThread(userID, username, formattedBody.Name, formattedBody.Description)
	json.NewEncoder(w).Encode(response)
}

// @Title Updates as a thread.
// @Description Updates a thread from a thread ID.
// @Param  threadID  path  int  true  "threadID of the thread in the path"
// @Param  thread  body  ThreadBody  true  "Info of the thread (name, description)"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  UpdateThreadResponse  "UpdateThreadResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource thread
// @Route /thread/{threadID} [put]
func UpdateThreadHandler(w http.ResponseWriter, r *http.Request) {
	// Get the data from the body
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	// Get the claims
	claims, _ := auth.GetClaims(r)

	// Format the data from the body
	var formattedBody ThreadBody
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		handleError(w, err)
		return
	}

	// Get the threadID from the path
	vars := mux.Vars(r)
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the response
	response, err := UpdateThread(uint(key), formattedBody.Name, formattedBody.Description, claims["username"].(string))
	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// @Title Delets as a thread.
// @Description Deletes a thread from a thread ID.
// @Param  threadID  path  int  true  "Thread ID from the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  DeleteThreadResponse  "DeleteThreadResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource thread
// @Route /thread/{threadID} [delete]
func DeleteThreadHandler(w http.ResponseWriter, r *http.Request) {
	// Get Claims
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)

	// Get threadID
	key, err := strconv.ParseUint(vars["threadID"], 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

	// Gets the response
	response, err := DeleteThread(uint(key), claims["username"].(string))
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
