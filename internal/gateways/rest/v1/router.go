package v1

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, ctrl *Controller) *gin.Engine {
	api := r.Group("/api")
	{
		api.GET("/get-all", ctrl.GetAll)

		private := api.Use(AuthMiddleware())
		{
			private.POST("/add", ctrl.Add)
		}
	}

	return r
}
