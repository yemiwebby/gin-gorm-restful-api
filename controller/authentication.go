package controller

import (
	"diary_api/model"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Helper that creates and sends token when a user logins & signup
	createSendToken(context, *savedUser, http.StatusCreated)
}

func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Helper that creates and sends token when a user logins & signup
	createSendToken(context, user, http.StatusCreated)
}

// JWT helper functions are for auth modules only,
// for better code readabilty / modularity, they should be in their scopes

func createSendToken(c *gin.Context, user model.User, status int) {
	jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	jwtExp, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	// Evaluated in days, should be modified to fit your preferences
	jwtDuration := time.Hour * 24 * time.Duration(jwtExp)
	jwtMaxAge := time.Now().Add(jwtDuration).Unix()

	jwt, err := generateJWT(user, jwtMaxAge, jwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", jwt, int(jwtDuration.Seconds()), "/", "localhost", false, true)

	c.JSON(status, gin.H{"token": jwt, "user": user})
}

func generateJWT(user model.User, maxAge int64, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		// using "eat" doesn't actually expire the token
		"exp": maxAge,
	})
	return token.SignedString(secret)
}
