package user

import (
	"backend-forum/auth"
	"backend-forum/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// @Title Gets as a user from a username.
// @Description Gets user info from a specific username.
// @Param  username  path  string  true  "username of the user in the path"
// @Success  200  object  UserResponse  "UserResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource user
// @Route /user/{username} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gets the username from the path
	vars := mux.Vars(r)
	key := vars["username"]
	response, err := GetUser(key)

	// Handle if any bad request error occurs
	if err != nil {
		utils.HandleErr(err)
		errResponse := ErrorResponse{Msg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// @Title Updates a user password.
// @Description Updates a user password, must be used with the user itself or a superuser.
// @Param  password  body  UpdateBody  true  "The new password in the body"
// @Success  200  object  UpdateResponse  "UpdateResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource user
// @Route /user/{username} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gets the data in the body
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	// Gets the claim in the token authorization header
	claims, _ := auth.GetClaims(r)

	// Format the data in the body
	var formattedBody UpdateBody
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	// Gets the username from the path
	vars := mux.Vars(r)
	key := vars["username"]

	response, err := UpdateUser(key, formattedBody.Password, claims["username"].(string))

	// Handle if any bad request error occurs
	if err != nil {
		utils.HandleErr(err)
		errResponse := ErrorResponse{Msg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)
	key := vars["username"]

	response := DeleteUser(key, claims["username"].(string))

	json.NewEncoder(w).Encode(response)
}
