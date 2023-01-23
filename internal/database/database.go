package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
}

// func GetDB() (*Database, error) {
// 	return &Database{}, nil
// }

func GetDB() (*sql.DB, error) {

	var db *sql.DB
	fmt.Println("Connecting to DB")
	db, err := sql.Open("mysql", "root:gya1ydxf@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		fmt.Println("Error connecting to DB:", err)
		return &sql.DB{}, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging to DB:", pingErr)
		return &sql.DB{}, pingErr
	}

	fmt.Println("Database connected!")

	return db, nil

}
