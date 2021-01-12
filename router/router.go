package router

import (
	"backend-forum/auth"
	"backend-forum/forum"
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
	authR.HandleFunc("/logout", auth.Middleware(user.LogoutHandler))

	forumR := r.PathPrefix("/forum").Subrouter()
	forumR.HandleFunc("/add", auth.Middleware(forum.AddForumHandler)).Methods("POST")
	forumR.HandleFunc("/{forumID}", forum.GetForumHandler).Methods("GET")
	forumR.HandleFunc("/{forumID}", auth.Middleware(forum.UpdateForumHandler)).Methods("POST")
	forumR.HandleFunc("/{forumID}", auth.Middleware(forum.DeleteForumHandler)).Methods("DELETE")

	// postR := r.PathPrefix("/post").Subrouter()

	log.Info("Server Running...")
	fmt.Println("Server Running...")
	log.Fatal(http.ListenAndServe(":8888", r))
}
