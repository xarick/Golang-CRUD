package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-crud/internal/models"
)

func (c *Controller) Add(ctx *gin.Context) {
	user := models.UserCrUp{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := c.serv.CRUDSer.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newUser})
}

func (c *Controller) GetUsers(ctx *gin.Context) {
	users, err := c.serv.CRUDSer.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *Controller) GetUser(ctx *gin.Context) {
	user, err := c.serv.CRUDSer.GetUser(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// if ID == "" {
// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
// 	return
// }
