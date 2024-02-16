package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliancampos/goopportunities/config"
)

func Initialize() {
	r := gin.Default()
	initializeRoutes(r)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	result := config.GetEnvValue("PORT")
	r.Run(":" + result)
}
