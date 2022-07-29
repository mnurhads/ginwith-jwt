package middlewares

import (
	"ginwith-jwt/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "Reqest failed, acces token not emty!"})
			context.Abort()
			return
		}
		err:= auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}