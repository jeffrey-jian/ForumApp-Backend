package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/comments"
	"github.com/CVWO/sample-go-app/internal/handlers/posts"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", GetUsers)
		r.Get("/posts", GetPosts)
		r.Get("/comments", GetComments)
		r.Post("/comments", CreateComment)
	}
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	response, _ := users.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPosts(w http.ResponseWriter, req *http.Request) {
	response, _ := posts.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetComments(w http.ResponseWriter, req *http.Request) {
	response, _ := comments.HandleList(w, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateComment(w http.ResponseWriter, req *http.Request) {
	comments.HandleCreateComment(w, req)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}
