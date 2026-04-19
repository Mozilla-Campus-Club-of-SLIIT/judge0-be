package handlers

import (
	"net/http"
	"strconv"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/logger"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/repositories"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/utils"
	"github.com/gin-gonic/gin"
)

func GetDSASubmissionResultsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	submissions, currentPage, totalPages, err := repositories.GetDSASubmissionResults(ctx, page, pageSize)
	if err != nil {
		logger.Log.Error("GetDSASubmissionResultsHandler: failed to get submission results", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"submissions": submissions,
		"currentPage": currentPage,
		"totalPages":  totalPages,
	})
}

func GetAllDSAChallengesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	challenges, currentPage, totalPages, err := repositories.GetAllDSAChallengesWithTestCases(ctx, page, pageSize)
	if err != nil {
		logger.Log.Error("GetAllDSAChallengesHandler: failed to get DSA challenges", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"challenges":  challenges,
		"currentPage": currentPage,
		"totalPages":  totalPages,
	})
}

func GetLiveLeaderboardAdminHandler(c *gin.Context) {
	ctx := c.Request.Context()
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	users, currentPage, totalPages, err := repositories.GetLiveLeaderboard(ctx, page, pageSize)
	if err != nil {
		logger.Log.Error("GetLiveLeaderboardAdminHandler: failed to get live leaderboard", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"currentPage": currentPage,
		"totalPages":  totalPages,
		"users":       users,
	})
}

func GetAllLinuxChallengesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	challenges, currentPage, totalPages, err := repositories.GetAllLinuxChallenges(ctx, page, pageSize)
	if err != nil {
		logger.Log.Error("GetAllLinuxChallengesHandler: failed to get Linux challenges", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"challenges":  challenges,
		"currentPage": currentPage,
		"totalPages":  totalPages,
	})
}

func UpdateChallengeStatusHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	statusID, err := strconv.Atoi(status)
	if err != nil || (statusID != 1 && statusID != 2) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 1 (inactive) or 2 (active)"})
		return
	}

	updated, err := repositories.UpdateChallengeStatus(ctx, id, statusID)
	if err != nil {
		logger.Log.Error("UpdateChallengeStatusHandler: failed to update challenge status", "id", id, "status_id", statusID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "challenge not found"})
		return
	}

	message := "Challenge status updated"
	if statusID == 1 {
		message = "Challenge deactivated"
	}
	if statusID == 2 {
		message = "Challenge activated"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"id":      id,
		"status":  statusID,
	})
}

func UpdateAllDSAChallengeStatusesHandler(c *gin.Context) {
	ctx := c.Request.Context()

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	statusID, err := strconv.Atoi(status)
	if err != nil || (statusID != 1 && statusID != 2) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 1 (inactive) or 2 (active)"})
		return
	}

	updatedCount, err := repositories.UpdateAllDSAChallengeStatuses(ctx, statusID)
	if err != nil {
		logger.Log.Error("UpdateAllDSAChallengeStatusesHandler: failed to update DSA challenge statuses", "status_id", statusID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := "All DSA challenge statuses updated"
	if statusID == 1 {
		message = "All DSA challenges deactivated"
	}
	if statusID == 2 {
		message = "All DSA challenges activated"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       message,
		"status":        statusID,
		"updated_count": updatedCount,
	})
}

func UpdateAllLinuxChallengeStatusesHandler(c *gin.Context) {
	ctx := c.Request.Context()

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	statusID, err := strconv.Atoi(status)
	if err != nil || (statusID != 1 && statusID != 2) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 1 (inactive) or 2 (active)"})
		return
	}

	updatedCount, err := repositories.UpdateAllLinuxChallengeStatuses(ctx, statusID)
	if err != nil {
		logger.Log.Error("UpdateAllLinuxChallengeStatusesHandler: failed to update Linux challenge statuses", "status_id", statusID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := "All Linux challenge statuses updated"
	if statusID == 1 {
		message = "All Linux challenges deactivated"
	}
	if statusID == 2 {
		message = "All Linux challenges activated"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       message,
		"status":        statusID,
		"updated_count": updatedCount,
	})
}

func DeleteChallengeHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	deleted, err := repositories.DeleteChallengeByID(ctx, id)
	if err != nil {
		logger.Log.Error("DeleteChallengeHandler: failed to delete challenge", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "challenge not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Challenge deleted",
		"id":      id,
	})
}

func GetJudge0SubmissionDetailsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	result, err := utils.GetJudge0SubmissionDetails(ctx, token, c.Request.URL.RawQuery)
	if err != nil {
		logger.Log.Error("GetJudge0SubmissionDetailsHandler: failed to fetch submission details", "token", token, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", result)
}

func ToggleLeaderboardFreezeHandler(c *gin.Context) {
	ctx := c.Request.Context()

	frozen, err := repositories.ToggleLeaderboardFreeze(ctx)
	if err != nil {
		logger.Log.Error("ToggleLeaderboardFreezeHandler: failed to toggle leaderboard freeze", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := "Leaderboard is now live"
	if frozen {
		message = "Leaderboard is now frozen"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"setting": "leaderboard_freeze",
		"value":   frozen,
	})
}
