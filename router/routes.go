package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliancampos/goopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
	}
}
