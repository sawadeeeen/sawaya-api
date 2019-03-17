package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/internal/apps/article-api/createArticle"
	"github.com/sawadeeeen/sawaya-api/internal/apps/article-api/getArticles"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", getArticles.GetArticles)
		v1.GET("/create", createArticle.CreateArticle)
		// v1.GET("/:id", getArticleById)
		// v1.POST("/:id", editArticle)
		// v1.POST("/", deleteArticle)
	}
	router.Run()
}
