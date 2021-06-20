package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
)

func init() {
	var err error

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	Client, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connected succesfully")
}
