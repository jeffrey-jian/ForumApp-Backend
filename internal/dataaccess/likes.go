package dataaccess

import (
	"database/sql"

	"github.com/CVWO/sample-go-app/internal/models"
)

func GetLikes(db *sql.DB, post_id string) ([]models.Like, error) {

	var results *sql.Rows
	var err error
	if post_id != "" {
		results, err = db.Query(`SELECT * FROM Likes WHERE post_id=` + post_id)
	} else {
		results, err = db.Query("SELECT * FROM Likes")
	}
	if err != nil {
		panic(err.Error())
	}
	likes := []models.Like{}

	for results.Next() {
		var like models.Like

		err = results.Scan(&like.ID, &like.User_ID, &like.Post_ID)
		if err != nil {
			panic(err.Error())
		}

		likes = append(likes, like)
	}
	return likes, nil
}
