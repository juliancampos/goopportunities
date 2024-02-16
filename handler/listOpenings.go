package handler

import "github.com/gin-gonic/gin"

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "showOpeningHandler",
	})
}
