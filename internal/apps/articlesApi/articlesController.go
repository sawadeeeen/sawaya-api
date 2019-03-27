package articlesApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	//"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/implements"
	"encoding/json"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	//	"strconv"
)

func CreateArticle(c *gin.Context) {

	c.Request.ParseForm()
	var article *model.Article
	c.BindJSON(&article)

	// can't
	//err := implements.Store(article)
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

	_, err = res.Exec(article.Title, article.Content, article.Published)
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
	//	var articles model.Articles
	var articles []model.Article
	for rows.Next() {
		var article model.Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Published, &article.Published_at); err != nil {
			panic(err.Error())
		}
		articles = append(articles, article)
	}
	jsonArticles, _ := json.Marshal(articles)
	c.String(http.StatusOK, string(jsonArticles))
}

func GetArticleById(c *gin.Context) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var article model.Article
	id := c.Param("id")
	err = db.QueryRow("SELECT * FROM article WHERE id = ?", id).Scan(&article.Id, &article.Title, &article.Content, &article.Published, &article.Published_at)
	if err != nil {
		panic(err.Error())
	}
	jsonArticle, _ := json.Marshal(article)
	c.String(http.StatusOK, string(jsonArticle))

}

func UpdateArticle(c *gin.Context) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	id := c.Param("id")

	c.Request.ParseForm()
	var article *model.Article
	c.BindJSON(&article)

	res, err := db.Prepare(`UPDATE article SET title=?,content=?,published=?,Published_at=now() WHERE id=?;`)
	if err != nil {
		panic(err.Error())
	}
	_, err = res.Exec(article.Title, article.Content, article.Published, id)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteArticle(c *gin.Context) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	id := c.Param("id")

	res, err := db.Prepare(`DELETE FROM article WHERE id = ? LIMIT 1`)
	if err != nil {
		err.Error()
	}
	_, err = res.Exec(id)
	if err != nil {
		err.Error()
	}
}
