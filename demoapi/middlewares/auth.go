package middlewares

import (
	"demoapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Please login to access this route.",
			},
		)
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Please login to access this route.",
			},
		)
		return
	}

	context.Set("userId", userId)
	context.Next()
}
