package user

import (
	"backend-forum/auth"
	"backend-forum/interfaces"
	"backend-forum/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["username"]
	response := GetUser(key)
	json.NewEncoder(w).Encode(response)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)
	claims, _ := auth.GetClaims(r)

	var formattedBody interfaces.User
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	vars := mux.Vars(r)
	key := vars["username"]

	response := UpdateUser(key, formattedBody.Password, claims["username"].(string))

	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth.GetClaims(r)
	vars := mux.Vars(r)
	key := vars["username"]

	response := DeleteUser(key, claims["username"].(string))

	json.NewEncoder(w).Encode(response)
}
