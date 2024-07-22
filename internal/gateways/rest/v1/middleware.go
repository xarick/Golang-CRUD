package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-crud/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		if len(headerAuth) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -101, "error": "login or password empty"})
			return
		}

		login, password, err := pkg.BasicAuthLogPass(headerAuth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -102, "error": err.Error()})
			return
		}

		if login == "admin" && password == "12345" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -103, "error": "login or password incorrect"})
			return
		}
	}
}
