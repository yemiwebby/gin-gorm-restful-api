package middleware

import (
	"diary_api/helper"
	"diary_api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := helper.ValidateJWT(context)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		userId := uint(claims["id"].(float64))

		user, err := model.FindUserById(userId)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusNotFound, gin.H{"error": err})
			context.Abort()
			return
		}

		// user can be accessed and modified if in other controllers if need be
		context.Set("user", &user)

		context.Next()
	}
}
