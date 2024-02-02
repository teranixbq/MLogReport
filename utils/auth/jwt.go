package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		data := token.Claims.(jwt.MapClaims)
		c.Set("user", gin.H{
			"id": data["id"],
			"role": data["role"],
		})

		c.Next()

		// if tokenString[0:7] != "Bearer " {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	return
		// }
	}
}

func CreateToken(id ,role string) (string, error) {
	config := viper.New()
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetString("JWT_SECRET")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	config := viper.New()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(c *gin.Context) (string, string, error) {

	user, exist := c.Get("user")
	if !exist {
		return "", "", errors.New("invalid token")
	}

	claims := user.(map[string]interface{})

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
