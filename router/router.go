package router

import (
	"backend-forum/auth"
	"backend-forum/post"
	"backend-forum/thread"
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

	threadR := r.PathPrefix("/thread").Subrouter()
	threadR.HandleFunc("/add", auth.Middleware(thread.AddThreadHandler)).Methods("POST")
	threadR.HandleFunc("/{threadID}", thread.GetThreadHandler).Methods("GET")
	threadR.HandleFunc("/{threadID}", auth.Middleware(thread.UpdateThreadHandler)).Methods("PUT")
	threadR.HandleFunc("/{threadID}", auth.Middleware(thread.DeleteThreadHandler)).Methods("DELETE")

	postR := r.PathPrefix("/thread/{threadID}/post").Subrouter()
	postR.HandleFunc("/add", auth.Middleware(post.AddPostHandler)).Methods("POST")
	postR.HandleFunc("/{postID}", post.GetPostHandler).Methods("GET")
	postR.HandleFunc("/{postID}", auth.Middleware(post.UpdatePostHandler)).Methods("PUT")

	log.Info("Server Running...")
	fmt.Println("Server Running...")
	log.Fatal(http.ListenAndServe(":8888", r))
}
