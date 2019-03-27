package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Init() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}
