package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"task-manager/pkg/jwtutil"

	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

func JWTAuthMiddleware(publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtutil.ValidateToken(tokenStr, publicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid subject claim"})
			return
		}

		c.Set(UserIDKey, userID)
		c.Next()
	}
}
