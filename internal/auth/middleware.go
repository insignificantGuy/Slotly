package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuth(tokens *Tokens) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing or invalid authorization header"})
			return
		}
		userID, err := tokens.ParseAccess(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid or expired token"})
			return
		}
		c.Set("user_id", userID)
		c.Request = c.Request.WithContext(WithUserID(c.Request.Context(), userID))
		c.Next()
	}
}

func CurrentUserID(c *gin.Context) (string, bool) {
	v, ok := c.Get("user_id")
	if !ok {
		return "", false
	}
	id, ok := v.(string)
	return id, ok && id != ""
}

func MustUserID(c *gin.Context) (string, bool) {
	id, ok := CurrentUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return "", false
	}
	return id, true
}
