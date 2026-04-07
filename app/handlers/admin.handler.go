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
