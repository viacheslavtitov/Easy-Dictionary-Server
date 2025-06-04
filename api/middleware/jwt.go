package middleware

import (
	internalenv "easy-dictionary-server/internalenv"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func JWTMiddleware(env *internalenv.Env, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			zap.S().Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		claims := token.Claims.(*Claims)
		if claims.Role != requiredRole {
			zap.S().Debugf("User role is not match. Now is %s , but required is %s", claims.Role, requiredRole)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient role"})
			return
		}
		c.Set("userUUID", claims.Subject)
		c.Set("user_role", claims.Role)
		c.Set("userID", claims.UserID)
		zap.S().Debugf("Set data into context %s, user id %d, user uuid %s", claims.Role, claims.UserID, claims.Subject)
		c.Next()
	}
}
