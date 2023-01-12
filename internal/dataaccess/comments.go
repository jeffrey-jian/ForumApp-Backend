package dataaccess

import (
	"database/sql"
	"fmt"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetComments(db *sql.DB, post_id string) ([]models.Comment, error) {

	var results *sql.Rows
	var err error
	if post_id != "" {
		results, err = db.Query(`SELECT Comments.id AS id, 
																		Comments.author_id AS author_id, 
																		Users.username AS author_username, 
																		Comments.date_created AS date_created, 
																		Comments.comment_text AS comment_text, 
																		Comments.post_id AS post_id 
																FROM Comments 
																JOIN Users ON Comments.author_id = Users.id AND Comments.post_id = ` + post_id)
	} else {
		results, err = db.Query(`SELECT Comments.id AS id, 
																		Comments.author_id AS author_id, 
																		Users.username AS author_username, 
																		Comments.date_created AS date_created, 
																		Comments.comment_text AS comment_text, 
																		Comments.post_id AS post_id 
															FROM Comments 
															JOIN Users 
															ON Comments.author_id = Users.id`)
	}
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(&results)
	comments := []models.Comment{}

	for results.Next() {
		var comment models.Comment

		err = results.Scan(&comment.ID, &comment.Author_ID, &comment.Author_Username, &comment.Date_created, &comment.Comment_text, &comment.Post_ID)
		if err != nil {
			panic(err.Error())
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
