package auth

import (
	"net/http"

	"github.com/Jeno7u/studybud/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr, err := c.Cookie("auth_token")
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing auth token"})
            return
        }

        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return config.Envs.JWTSecret, nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        // Extract claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "bad claims"})
            return
        }

        // Pass user data to handlers
        c.Set("user_id", int(claims["user_id"].(float64)))

        c.Next()
    }
}