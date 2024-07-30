package v1

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, ctrl *Controller) *gin.Engine {
	api := r.Group("/api")
	{
		api.GET("/users", ctrl.GetUsers)
		api.GET("/user/:id", ctrl.GetUser)

		private := api.Use(AuthMiddleware())
		{
			private.POST("/add", ctrl.Add)
			private.PUT("/update/:id", ctrl.Update)
			private.DELETE("/delete/:id", ctrl.Delete)
		}
	}

	return r
}
