package middleware

import (
	"net/http"
	"scrap-invoice-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware - Validasi JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RoleMiddleware - Validasi role user (RBAC)
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		userRole := role.(string)

		// Cek apakah role user ada di daftar allowedRoles
		allowed := false
		for _, r := range allowedRoles {
			if r == userRole {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"error":         "You don't have permission to access this resource",
				"your_role":     userRole,
				"allowed_roles": allowedRoles,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
