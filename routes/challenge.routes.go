package routes

import (
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/handlers"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ChallengeRoutes(r *gin.RouterGroup) {
	challenge := r.Group("/challenge")

	challenge.GET("/:id", handlers.GetChallengeByID)
	challenge.GET("/get", handlers.GetChallenges)
	challenge.GET("/test", middleware.RoleRequiredMiddleware([]string{"admin", "user"}), handlers.TestChallenge)
}
