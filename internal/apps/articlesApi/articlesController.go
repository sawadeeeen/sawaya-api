package articlesApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	//	"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/implements"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreateArticle(c *gin.Context) {

	c.Request.ParseForm()
	inputData := model.Articles{
		//id := c.Request.Form["id"],
		//	Title:        c.Request.Form["title"],
		//	Content:      c.Request.Form["content"],
		//	Published:    c.Request.Form["published"],
		//	Published_at: c.Request.Form["published_at"],
		Title:        "gafa",
		Content:      "gafa",
		Published:    true,
		Published_at: "1111",
	}

	// 	implements.Store(inputData)
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}

	res, err := db.Prepare(
		`insert into article(title, content, published, published_at) values(?,?,?,?)`,
	)
	//_, err = db.Exec()
	if err != nil {
		panic(err.Error())
	}

	_, err = res.Exec(inputData.Title, inputData.Content, inputData.Published, inputData.Published_at)
	//	_, err = res.Exec(Articles.Title, Articles.Content, Articles.Published, Articles.Published_at)
	//	return err

	defer db.Close()
	//  c.String(http.StatusOK, "hello")
	// c.Request.ParseForm()
	// fmt.Println(c.Request.Form["id"])
}

func GetArticles(c *gin.Context) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	//res, err := db.Prepare(
	//	`select * from article`,
	//	)
	var title string
	if err := db.QueryRow("SELECT  title FROM article WHERE id = 1 LIMIT 1").Scan(&title); err != nil {
		//log.Fatal(err)
		//	panic(err.Error())
	}

	//_, err = db.Exec()
	//fmt.Println(id)

	c.String(http.StatusOK, title)
}
