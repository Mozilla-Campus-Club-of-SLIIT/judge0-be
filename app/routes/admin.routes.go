package routes

import (
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/handlers"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	challenge := r.Group("/admin")
	{
		challenge.GET("/submissions/dsa", middlewares.AuthMiddleware("Codenight host"), handlers.GetDSASubmissionResultsHandler)
		challenge.GET("/dsa/challenges", middlewares.AuthMiddleware("Codenight host"), handlers.GetAllDSAChallengesHandler)
		challenge.PATCH("/challenges/:id/:status", middlewares.AuthMiddleware("Codenight host"), handlers.UpdateChallengeStatusHandler)
		challenge.PATCH("/dsa/challenges/:status", middlewares.AuthMiddleware("Codenight host"), handlers.UpdateAllDSAChallengeStatusesHandler)
		challenge.GET("/submissions/dsa/:token/details", middlewares.AuthMiddleware("Codenight host"), handlers.GetJudge0SubmissionDetailsHandler)
	}
}
