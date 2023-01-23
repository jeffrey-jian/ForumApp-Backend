package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	da "github.com/CVWO/sample-go-app/internal/dataaccess"
	"github.com/CVWO/sample-go-app/internal/database"

	"github.com/CVWO/sample-go-app/internal/models"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

const (
	ListPosts  = "posts.HandleList"
	CreatePost = "posts.HandleCreatePost"
	EditPost   = "posts.HandleEditPost"
	DeletePost = "posts.HandleDeletePost"

	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
	ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListPosts))
	}

	id := r.URL.Query().Get("id")
	filter := r.URL.Query().Get("filter")
	searchTerm := r.URL.Query().Get("searchTerm")
	author := r.URL.Query().Get("author")
	likedBy := r.URL.Query().Get("likedBy")

	fmt.Println("Running da.GetPosts()...")
	posts, err := da.GetPosts(db, id, filter, searchTerm, author, likedBy)
	if err != nil {
		fmt.Println("Error while running da.GetPosts()")
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data, err := json.Marshal(posts)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListPosts))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	// fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, CreatePost))
	}

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare(
		`INSERT INTO Posts (author_id, category, title, post_text) 
			VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, er := query.Exec(post.Author_ID, post.Category, post.Title, post.Post_text)
	if er != nil {
		return nil, er
	}
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	return nil, nil
}

func HandleEditPost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, EditPost))
	}

	var post models.Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("UPDATE Posts SET category=?, title=?, post_text=? WHERE id=?")
	if err != nil {
		return nil, err
	}

	_, er := query.Exec(post.Category, post.Title, post.Post_text, id)
	if er != nil {
		return nil, er
	}
	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
	return nil, nil
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, DeletePost))
	}
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		return nil, err
	}
	_, er := query.Exec(id)
	if er != nil {
		return nil, er
	}
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})

	return nil, nil
}
