package dataaccess

import (
	"database/sql"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetComments(db *sql.DB, post_id string) ([]models.Comment, error) {

	var results *sql.Rows
	var err error

	if post_id != "" {
		results, err = db.Query("SELECT * FROM comments WHERE post_id = " + post_id)
	} else {
		results, err = db.Query("SELECT * FROM comments")
	}

	if err != nil {
		panic(err.Error())
	}

	comments := []models.Comment{}

	for results.Next() {
		var comment models.Comment

		err = results.Scan(&comment.ID, &comment.Author_ID, &comment.Date_created, &comment.Comment_text, &comment.Post_ID)
		if err != nil {
			panic(err.Error())
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
