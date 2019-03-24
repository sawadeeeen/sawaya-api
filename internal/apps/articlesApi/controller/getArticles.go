package getArticles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
