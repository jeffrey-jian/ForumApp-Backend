package dataaccess

import (
	"database/sql"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetPosts(db *sql.DB, id string) ([]models.Post, error) {

	var results *sql.Rows
	var err error

	if id != "" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id AND Post.id = ` + id)
	} else {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id ORDER BY date_created DESC`)
	}

	if err != nil {
		panic(err.Error())
	}

	posts := []models.Post{}

	for results.Next() {
		var post models.Post

		err = results.Scan(&post.ID, &post.Author_ID, &post.Author_Username, &post.Category, &post.Date_created, &post.Title, &post.Post_text)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	return posts, nil

}
