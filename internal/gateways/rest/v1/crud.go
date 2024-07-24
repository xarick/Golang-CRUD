package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Get All"})
}

func (c *Controller) Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "Add struct"})
}
