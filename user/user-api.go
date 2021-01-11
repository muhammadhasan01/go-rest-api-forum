package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"backend-forum/auth"
	"backend-forum/utils"
)

type loginBody struct {
	Username string
	Password string
}

type registerBody struct {
	Username string
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	var formattedBody loginBody
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)
	response := Login(formattedBody.Username, formattedBody.Password)

	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	var formattedBody registerBody
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	response := Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := auth.GetClaims(r)
	fmt.Println(claims)
	w.Write([]byte(`{"message": "logout called"}`))
}
