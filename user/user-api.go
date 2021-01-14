package user

import (
	"backend-forum/auth"
	"backend-forum/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// @Title Gets as a user.
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
// @Param  username  path  string  true  "username of the user in the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
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

// @Title Deletes a user.
// @Description Delete a user from the username path.
// @Param  username  path  string  true  "username of the user in the path"
// @Param  token  header  string  true  "JWT Token received when logged in"
// @Success  200  object  DeleteResponse  "DeleteResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource user
// @Route /user/{username} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gets claims
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)

	// Gets the username from the path
	key := vars["username"]

	response, err := DeleteUser(key, claims["username"].(string))

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
