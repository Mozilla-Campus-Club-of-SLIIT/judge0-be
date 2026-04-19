package routes

import (
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/handlers"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/middlewares"
	"github.com/gin-gonic/gin"
)

const adminRole = "Codenight host"

func AdminRoutes(r *gin.RouterGroup) {
	challenge := r.Group("/admin")
	{
		challenge.GET("/submissions/dsa", middlewares.AuthMiddleware(adminRole), handlers.GetDSASubmissionResultsHandler)
		challenge.GET("/dsa/challenges", middlewares.AuthMiddleware(adminRole), handlers.GetAllDSAChallengesHandler)
		challenge.GET("/linux/challenges", middlewares.AuthMiddleware(adminRole), handlers.GetAllLinuxChallengesHandler)
		challenge.PATCH("/leaderboard/toggle", middlewares.AuthMiddleware(adminRole), handlers.ToggleLeaderboardFreezeHandler)
		challenge.PATCH("/challenges/:id/:status", middlewares.AuthMiddleware(adminRole), handlers.UpdateChallengeStatusHandler)
		challenge.DELETE("/challenges/:id", middlewares.AuthMiddleware(adminRole), handlers.DeleteChallengeHandler)
		challenge.PATCH("/dsa/challenges/:status", middlewares.AuthMiddleware(adminRole), handlers.UpdateAllDSAChallengeStatusesHandler)
		challenge.PATCH("/linux/challenges/:status", middlewares.AuthMiddleware(adminRole), handlers.UpdateAllLinuxChallengeStatusesHandler)
		challenge.GET("/submissions/dsa/:token/details", middlewares.AuthMiddleware(adminRole), handlers.GetJudge0SubmissionDetailsHandler)
	}
}
