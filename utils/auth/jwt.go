package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	config := viper.New()
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {	
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"].(string))
			c.Set("role", claims["role"].(string))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}

func CreateToken(id string, role string) (string, error) {
	config := viper.New()
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetString("JWT_SECRET")))
}

func ExtractToken(c *gin.Context) (string, string, error) {
	user, exists := c.Get("user")
	if !exists {
		return "", "", errors.New("invalid token")
	}

	token, ok := user.(*jwt.Token)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	return id, role, nil
}
