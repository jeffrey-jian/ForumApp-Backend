package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
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
	// for localhost use
	// db, err := sql.Open("mysql", "root:gya1ydxf@tcp(127.0.0.1:3306)/forum")

	// hardcoded address for clearDB connection
	// db, err := sql.Open("mysql", "b83f3be08aeca3:acf4cd99@tcp(us-cdbr-east-06.cleardb.net)/heroku_375321b59849bf2")

	cfg := mysql.Config{
		User:                 "b83f3be08aeca3",
		Passwd:               "acf4cd99",
		Net:                  "tcp",
		Addr:                 "us-cdbr-east-06.cleardb.net",
		DBName:               "heroku_375321b59849bf2",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

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
