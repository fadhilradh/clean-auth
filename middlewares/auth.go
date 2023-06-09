package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Access token required",
		})
		c.Abort()
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Access token is expired",
			})
			c.Abort()
		}

		if claims["role"] != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Insufficient role",
			})
			c.Abort()
		}
		c.Next()
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Access token is invalid",
		})
		c.Abort()
	}

}

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Access token required",
		})
		c.Abort()
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Access token is expired",
			})
			c.Abort()
		}

		role := claims["role"]
		sufficientRole := role == "user" || role == "admin"
		if !sufficientRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Insufficient role",
			})
			c.Abort()
		}
		c.Next()
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Access token is invalid",
		})
		c.Abort()
	}

}
