package router

import (
	// "backend-forum/auth"
	"backend-forum/user"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func StartAPI() {
	r := mux.NewRouter()
	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/login", user.LoginHandler).Methods("POST")
	authR.HandleFunc("/register", user.RegisterHandler).Methods("POST")
	// authR.HandleFunc("/logout", users.LogoutHandler)
	// http.Handle("/auth/logout", auth.Middleware(authR))

	// forumR := r.PathPrefix("/forum").Subrouter()

	// postR := r.PathPrefix("/post").Subrouter()

	log.Info("Server running...")
	// TODO: DELETE THIS LATER
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8888", r))
}
