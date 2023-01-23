package dataaccess

import (
	"database/sql"
	"fmt"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetPosts(db *sql.DB, id string, filter string, searchTerm string, author string, likedBy string) ([]models.Post, error) {

	var results *sql.Rows
	var err error

	if id != "" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id AND Posts.id = ` + id)
	} else if searchTerm != "" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id 
																AND (
																	Posts.title LIKE '%` + searchTerm + `%' OR
																	Posts.post_text LIKE '%` + searchTerm + `%'
																) ORDER BY date_created DESC`)
	} else if author != "" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id AND Posts.author_id = ` + author)
	} else if likedBy != "" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id 
																JOIN Likes ON Likes.post_id = Posts.id AND Likes.user_id = ` + likedBy)
	} else if filter != "All" && filter != "all" {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id AND Posts.Category = '` + filter + `' ORDER BY date_created DESC`)
	} else {
		results, err = db.Query(`SELECT Posts.id AS id,
																		Posts.author_id AS author_id,
																		Users.username AS author_username,
																		Users.avatarColor AS author_avatarColor,
																		Posts.Category AS category,
																		Posts.date_created AS date_created,
																		Posts.title AS title,
																		Posts.post_text AS post_text
																FROM Posts
																JOIN Users ON Posts.author_id = Users.id ORDER BY date_created DESC`)
	}

	if err != nil {
		fmt.Println("Error in dataacess/posts.go - db.Query", err)
		panic(err.Error())
	}

	posts := []models.Post{}

	for results.Next() {
		var post models.Post

		err = results.Scan(&post.ID, &post.Author_ID, &post.Author_Username, &post.Author_AvatarColor, &post.Category, &post.Date_created, &post.Title, &post.Post_text)
		if err != nil {
			fmt.Println("Error in dataacess/posts.go - results.Scan", err)
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	return posts, nil

}
