package likes

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
	ListLikes  = "likes.HandleList"
	AddLike    = "likes.HandleAddLike"
	DeleteLike = "likes.HandleDeleteLike"

	SuccessfulListCommentsMessage = "Successfully listed likes"
	ErrRetrieveDatabase           = "Failed to retrieve database in %s"
	ErrRetrieveLikes              = "Failed to retrieve likes in %s"
	ErrEncodeView                 = "Failed to retrieve likes in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListLikes))
	}

	post_id := r.URL.Query().Get("post_id")

	posts, err := da.GetLikes(db, post_id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveLikes, ListLikes))
	}

	data, err := json.Marshal(posts)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListLikes))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	// fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HandleAddLike(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, AddLike))
	}

	var like models.Like
	json.NewDecoder(r.Body).Decode(&like)

	query, err := db.Prepare(`INSERT INTO Likes (user_id, post_id) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}

	_, er := query.Exec(like.User_ID, like.Post_ID)

	if er != nil {
		return nil, er
	}
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully liked"})

	return nil, nil
}

func HandleDeleteLike(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, DeleteLike))
	}

	id := chi.URLParam(r, "id")

	query, err := db.Prepare("DELETE FROM likes WHERE id=?")
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
