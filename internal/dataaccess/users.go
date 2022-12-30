package users

import (
	"database/sql"

	"github.com/CVWO/sample-go-app/internal/models"
)

// func List(db *sql.DB) ([]models.User, error) {
// 	users := []models.User{
// 		{
// 			ID:   1,
// 			Name: "CVWO",
// 		},
// 	}
// 	return users, nil
// }

func List(db *sql.DB, id string) ([]models.User, error) {

	var results *sql.Rows
	var err error

	if id != "" {
		results, err = db.Query("SELECT * FROM users WHERE id = " + id)
	} else {
		results, err = db.Query("SELECT * FROM users")
	}

	if err != nil {
		panic(err.Error())
	}

	users := []models.User{}

	for results.Next() {
		var user models.User

		err = results.Scan(&user.ID, &user.Username)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
