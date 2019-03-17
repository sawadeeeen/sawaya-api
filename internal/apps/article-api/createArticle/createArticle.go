package createArticle

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func CreateArticle(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`insert into article(title, content, published, published_at) values("faffa","hello",0,now());`)
	if err != nil {
		panic(err.Error())
	}
	c.String(http.StatusOK, "hello")

}
