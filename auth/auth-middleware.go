package auth

import (
	"backend-forum/interfaces"
	"backend-forum/utils"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	jwt "github.com/dgrijalva/jwt-go"
)

const CtxKey = "auth-token"

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := ExtractToken(r)
			if err != nil {
				utils.HandleErr(err)
				fmt.Fprintf(w, err.Error())
				return
			}
			if !token.Valid {
				utils.HandleErr(errors.New("Token not valid"))
				fmt.Fprintf(w, "Error: Token not valid")
				return
			}

			if !checkTokenInDB(token.Raw) {
				utils.HandleErr(errors.New("Token not found in whitelist"))
				fmt.Fprintf(w, "Token not found in whitelist")
				return
			}

			next(w, r)
			return
		}
		log.Warn("User not authorized")
		fmt.Fprintf(w, "Not Authorized")
	})
}

func ExtractToken(r *http.Request) (*jwt.Token, error) {
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.HandleErr(errors.New("There was an error with the token"))
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(utils.GetEnv("API_SECRET")), nil
	})
	return token, err
}

func GetClaims(r *http.Request) (jwt.MapClaims, bool) {
	tokenStr := r.Header["Token"][0]

	hmacSecretString := utils.GetEnv("API_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["user_id"] = uint(claims["user_id"].(float64))
		claims["username"] = string(claims["username"].(string))
		return claims, true
	} else {
		log.Warn("Invalid JWT Token")
		return nil, false
	}
}

func GetToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(utils.GetEnv("SIGN_METHOD")), tokenContent)
	token, err := jwtToken.SignedString([]byte(utils.GetEnv("API_SECRET")))
	utils.HandleErr(err)

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
