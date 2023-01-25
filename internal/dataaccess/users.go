package dataaccess

import (
	"database/sql"
	"fmt"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func GetUsers(username string, avatarColor string) ([]models.User, error) {

	var db = database.DB
	var results *sql.Rows
	var err error

	users := []models.User{}
	if username != "" {
		var user models.User
		err = db.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&user.ID, &user.Username, &user.AvatarColor)
		switch {
		case err == sql.ErrNoRows:
			fmt.Printf("creating new user with username %s\n", username)
			results, err = db.Query("INSERT INTO Users (username, avatarColor) VALUES ('" + username + "', '" + avatarColor + "')")
			err = db.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&user.ID, &user.Username, &user.AvatarColor)
			fmt.Printf("username is %s\n", username)
			users = append(users, user)
		case err != nil:
			panic(err.Error())
		default:
			fmt.Printf("username is %s\n", username)
			users = append(users, user)
		}

	} else {
		results, err = db.Query("SELECT * FROM users")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var user models.User

			err = results.Scan(&user.ID, &user.Username, &user.AvatarColor)
			if err != nil {
				panic(err.Error())
			}
			users = append(users, user)
		}
	}

	return users, nil
}
