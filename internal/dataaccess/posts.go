package dataaccess

import (
	"database/sql"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetPosts(db *sql.DB, id string) ([]models.Post, error) {

	var results *sql.Rows
	var err error

	if id != "" {
		results, err = db.Query("SELECT * FROM posts WHERE id = " + id)
	} else {
		results, err = db.Query("SELECT * FROM posts")
	}

	if err != nil {
		panic(err.Error())
	}

	posts := []models.Post{}

	for results.Next() {
		var post models.Post

		err = results.Scan(&post.ID, &post.Author_ID, &post.Category, &post.Date_created, &post.Title, &post.Post_text)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	return posts, nil

}
