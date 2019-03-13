package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/internal/apps/article-api/getArticles"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", getArticles.GetArticles)
		// router.POST("/create", createArticle)
		// router.GET("/:id", getArticleById)
		// router.POST("/:id", editArticle)
		// router.POST("/", deleteArticle)
	}
	router.Run()
}

//func getArticles(c *gin.Context) {
//	c.String(http.StatusOK, "hello")
//}
