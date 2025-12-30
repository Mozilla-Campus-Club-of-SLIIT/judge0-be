package routes

import (
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/:id", handlers.GetUser)
	user.POST("", handlers.CreateUser)
}
