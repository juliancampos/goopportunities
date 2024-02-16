package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "showOpeningHandler",
	})
}
