package articlesApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	//"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/implements"
	"encoding/json"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreateArticle(c *gin.Context) {

	c.Request.ParseForm()
	var articles *model.Articles
	c.BindJSON(&articles)

	// can't
	//err := implements.Store(articles)
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}

	res, err := db.Prepare(
		`insert into article(title, content, published, published_at) values(?,?,?,now())`,
	)
	if err != nil {
		panic(err.Error())
	}

	_, err = res.Exec(articles.Title, articles.Content, articles.Published)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func GetArticles(c *gin.Context) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM article ")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//var articles []model.Articles
	for rows.Next() {
		var article model.Articles
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Published, &article.Published_at); err != nil {
			panic(err.Error())
		}
		// articleForAdding := model.Articles{article}
		//articles = append(articles, article)
		//	fmt.Println(json.Marshal(articles))
		b, _ := json.Marshal(article)
		c.String(http.StatusOK, string(b))
	}

}
