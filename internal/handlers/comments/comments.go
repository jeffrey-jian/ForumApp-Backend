package comments

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
	ListComments = "comments.HandleList"

	SuccessfulListCommentsMessage = "Successfully listed comments"
	ErrRetrieveDatabase           = "Failed to retrieve database in %s"
	ErrRetrieveComments           = "Failed to retrieve comments in %s"
	ErrEncodeView                 = "Failed to retrieve comments in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
	}

	post_id := r.URL.Query().Get("post_id")

	posts, err := da.GetComments(db, post_id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, ListComments))
	}

	data, err := json.Marshal(posts)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
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

func HandleCreateComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
	}

	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	fmt.Println("handling creating comment...")

	query, err := db.Prepare(
		`INSERT INTO Comments (author_id, comment_text, post_id) 
				VALUES (?, ?, ?)`)

	if err != nil {
		return nil, err
	}
	_, er := query.Exec(comment.Author_ID, comment.Comment_text, comment.Post_ID)

	if er != nil {
		return nil, er
	}
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})

	return nil, nil

}

func HandleEditComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
	}

	var comment models.Comment
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&comment)

	query, err := db.Prepare("UPDATE Comments SET comment_text=? WHERE id=?")
	if err != nil {
		return nil, err
	}

	_, er := query.Exec(comment.Comment_text, id)
	if er != nil {
		return nil, er
	}

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "edited successfully"})

	return nil, nil

}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
	}

	id := chi.URLParam(r, "id")

	query, err := db.Prepare("DELETE FROM comments WHERE id=?")
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
