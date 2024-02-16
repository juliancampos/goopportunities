package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliancampos/goopportunities/schemas"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error listing openings: %v", err.Error)
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
