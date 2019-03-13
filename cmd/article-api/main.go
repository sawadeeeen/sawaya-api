package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		//	router.GET("/", getArticles)
		//	router.POST("/create", createArticle)
		//	router.GET("/:id", getArticleById)
		//	router.POST("/:id", editArticle)
		//	router.POST("/", deleteArticle)
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello")
		})
	}
	router.Run()
}
