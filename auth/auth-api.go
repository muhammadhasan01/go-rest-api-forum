package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"backend-forum/utils"
)

// @Title Login as a user.
// @Description Handling a user to login, and creates a JWT Token for the user.
// @Param  user  body  LoginBody  true  "Info of a user (username and password)."
// @Success  202  object  LoginResponse  "LoginResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource auth
// @Route /auth/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Gets the data in the request body
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	// Format the data in the request body
	var formattedBody LoginBody
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)
	response, err := Login(formattedBody.Username, formattedBody.Password)

	// Handle if any bad request error occurs
	if err != nil {
		utils.HandleErr(err)
		errResponse := ErrorResponse{Msg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	// Set status to Accepted
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	var formattedBody RegisterBody
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)

	response := Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	claims, _ := GetClaims(r)
	token := r.Header["Token"][0]
	response := Logout(claims["user_id"].(uint), token, claims["username"].(string))
	json.NewEncoder(w).Encode(response)
}
