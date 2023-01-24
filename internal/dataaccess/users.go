package dataaccess

import (
	"database/sql"
	"fmt"

	"github.com/CVWO/sample-go-app/internal/database"
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
		// results, err = db.Query("INSERT INTO Users (username) VALUES ('" + username + "')")
		// if err != nil {
		// 	if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
		// 		if driverErr.Number == 1062 {
		// 			// Handle the duplicate entry error
		// 			fmt.Println("======= user already created =========")
		// 			results, err = db.Query("SELECT * FROM users WHERE username='" + username + "'")
		// 			if err != nil {
		// 				panic(err.Error())
		// 			}
		// 		}
		// 	}
		// }

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
