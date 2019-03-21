package dao

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func create() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

}
