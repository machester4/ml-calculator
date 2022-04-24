package ginhandlers

import (
	"net/http"

	"calculator.com/internal/engine"
	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context, e engine.ServiceEngine) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
