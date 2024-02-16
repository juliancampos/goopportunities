package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "showOpeningHandler",
	})
}
