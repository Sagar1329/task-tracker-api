package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"


	"github.com/Sagar1329/task-tracker-api/internal/auth"
)


func AuthMiddleware(jwtSecret string) gin.HandlerFunc{

	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"authorization header is required",
			})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "

		if !strings.HasPrefix(authHeader,bearerPrefix){
			c.JSON(http.StatusUnauthorized,gin.H{
				"error": "invalid authorization header",
			})
			c.Abort()
			return

		}


		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		claims,err := auth.ValidateToken(
			tokenString,
			jwtSecret,
		)


		if err != nil{
				c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}
	c.Set("userID", claims.UserID)

		c.Next()

	}
}