package middleware

import (
	"net/http"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/config"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/utils"
	"github.com/gin-gonic/gin"
)

func RoleRequiredMiddleware(requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := utils.GetBearerToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(tokenStr, []byte(config.Get().SecretKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			c.Abort()
			return
		}

		roles, err := utils.GetRolesFromClaims(claims)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !utils.HasRequiredRole(roles, requiredRoles) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}
