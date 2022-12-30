package comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/api"
	da "github.com/CVWO/sample-go-app/internal/dataaccess"
	"github.com/CVWO/sample-go-app/internal/database"

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
