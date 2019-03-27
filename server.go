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
		v1.POST("/", articlesApi.CreateArticle)
		v1.GET("/:id", articlesApi.GetArticleById)
		v1.PUT("/:id", articlesApi.UpdateArticle)
		v1.DELETE("/:id", articlesApi.DeleteArticle)
	}
	router.Run()
}
