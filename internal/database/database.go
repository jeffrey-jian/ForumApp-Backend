package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetDB() (*sql.DB, error) {

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
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to DB:", err)
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error pinging to DB:", err)
		return nil, err
	}

	return DB, nil

}
