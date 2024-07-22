package v1

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, ctrl *Controller) *gin.Engine {
	api := r.Group("/api").Use(AuthMiddleware())
	{
		api.GET("/get-all", ctrl.GetAllController)
	}
	return r
}
