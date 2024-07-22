package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetAllController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Get All"})
}
