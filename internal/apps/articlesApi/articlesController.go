package articlesApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/implements"
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

	implements.Store(inputData)

	//  c.String(http.StatusOK, "hello")
	// c.Request.ParseForm()
	// fmt.Println(c.Request.Form["id"])
}

func GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
