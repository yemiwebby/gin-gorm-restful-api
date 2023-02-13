package middleware

import (
	"diary_api/model"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := validateJWT(context)
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

		// user can be accessed and modified in other controllers if need be
		context.Set("user", &user)

		context.Next()
	}
}

// JWT helper functions are for auth modules only,
// for better code readabilty / modularity, they should be in their scopes

func validateJWT(c *gin.Context) (*jwt.Token, error) {
	reqToken := getTokenFromRequest(c)

	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))

		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return token, nil
	}
	return nil, errors.New("invalid token provided")
}

func getTokenFromRequest(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	ht := strings.Split(token, " ")
	if len(ht) == 2 {
		fmt.Println("hToken : ", ht[1])
		return ht[1]
	}

	ct, err := c.Cookie("access_token")
	if err != nil {
		fmt.Println("cToken : ")
		return ct
	}
	return ""
}
