package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sawadeeeen/sawaya-api/internal/apps/articlesApi"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", articlesApi.GetArticles)
		v1.POST("/create", articlesApi.CreateArticle)
		// v1.GET("/:id", getArticleById)
		// v1.POST("/:id", editArticle)
		// v1.POST("/", deleteArticle)
	}
	router.Run()
}
