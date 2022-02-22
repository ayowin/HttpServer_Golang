package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Interactor *sql.DB
var InteractorError error

func init() {
	host := "127.0.0.1"
	port := "3306"
	database := "http-server-golang"
	username := "root"
	password := "123456"

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		database)

	Interactor, InteractorError = sql.Open("mysql", url)

	/* data source config check */
	if InteractorError != nil {
		panic("open failed with: " + InteractorError.Error())
	}

	/* connect check */
	InteractorError = Interactor.Ping()
	if InteractorError != nil {
		panic("connect failed with: " + InteractorError.Error())
	}

	/* max connection */
	Interactor.SetMaxOpenConns(100)
	/* max idle connection */
	Interactor.SetMaxIdleConns(20)
	/* connection max lifetime */
	Interactor.SetConnMaxLifetime(100 * time.Second)
}
