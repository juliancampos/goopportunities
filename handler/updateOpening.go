package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "showOpeningHandler",
	})
}
