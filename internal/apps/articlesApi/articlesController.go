package articlesApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	//"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/implements"
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
