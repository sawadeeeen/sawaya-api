package articlesApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(c *gin.Context) {
	//  db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	//  if err != nil {
	//      panic(err.Error())
	//  }
	//  defer db.Close()

	//  _, err = db.Exec(`insert into article(title, content, published, published_at) values("faffa","hello",0,now());`)
	//  if err != nil {
	//      panic(err.Error())
	//  }
	//  c.String(http.StatusOK, "hello")
	c.Request.ParseForm()
	fmt.Println(c.Request.Form["id"])
}

func GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
