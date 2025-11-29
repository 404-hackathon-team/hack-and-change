package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jeno7u/studybud/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Read JWT from cookie
		tokenString, err := c.Cookie("auth_token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth_token cookie"})
			c.Abort()
			return
		}

		// fetch secret at runtime (ensures env/config is loaded)
		secret := config.Envs.JWTSecret

		// Parse the token and ensure HMAC signing method
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		// Try common claim keys and types for user id
		var userID int
		found := false
		keys := []string{"user_id", "userID", "userId", "userid", "sub"}
		for _, k := range keys {
			if raw, ok := claims[k]; ok && raw != nil {
				switch v := raw.(type) {
				case float64:
					userID = int(v)
					found = true
				case string:
					if n, err := strconv.Atoi(v); err == nil {
						userID = n
						found = true
					}
				case json.Number:
					if n64, err := v.Int64(); err == nil {
						userID = int(n64)
						found = true
					}
				}
				if found {
					break
				}
			}
		}

		if !found {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token claims"})
			c.Abort()
			return
		}

		// Put user id into Gin context (as int, since handlers expect GetInt)
		c.Set("user_id", userID)

		c.Next()
	}
}
