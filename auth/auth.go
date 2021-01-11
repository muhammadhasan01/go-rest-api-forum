package auth

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(utils.GetEnv("SIGN_METHOD")), tokenContent)
	token, err := jwtToken.SignedString([]byte(utils.GetEnv("API_SECRET")))
	utils.HandleErr(err)

	// TODO: DELETE THIS LATER
	fmt.Println("The token for", user, "is", token)

	db := utils.ConnectDB()
	defer db.Close()
	auth := &interfaces.Auth{UserID: user.ID, Token: token}
	db.Create(&auth)

	return token
}

func checkTokenInDB(token string) bool {
	db := utils.ConnectDB()
	defer db.Close()

	auth := &interfaces.Auth{}
	return !db.Where("token = ? ", token).First(&auth).RecordNotFound()
}

func Middleware(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					utils.HandleErr(errors.New("There was an error with the token"))
					return nil, fmt.Errorf("There was an error")
				}
				return utils.GetEnv("API_SECRET"), nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if !token.Valid {
				utils.HandleErr(errors.New("Token not valid"))
				fmt.Fprintf(w, "Error: Token not valid")
			}

			// TODO: DELETE THIS LATER
			fmt.Println("THIS IS THE TOKEN => ", token.Raw)

			checkTokenInDB(token.Raw)

			endpoint(w, r)
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
