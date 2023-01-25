package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// getDB establishes connection with database, info from env
func GetDB() (*sql.DB, error) {

	fmt.Println("Connecting to DB")

	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	net := os.Getenv("DBNET")
	address := os.Getenv("DBADDRESS")
	dbName := os.Getenv("DBNAME")

	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  net,
		Addr:                 address,
		DBName:               dbName,
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
