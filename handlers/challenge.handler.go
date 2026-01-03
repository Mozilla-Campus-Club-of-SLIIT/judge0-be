package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/repository"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/types"
	"github.com/gin-gonic/gin"
)

func GetChallengeByID(c *gin.Context) {
	id := c.Param("id")
	challenge, err := repository.GetChallengeByID(c.Request.Context(), id)

	if err != nil {
		log.Printf("Internal error while fetching challenge ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "challenge may not exist"})
		return
	}
	if challenge == nil {
		log.Printf("Challenge not found with ID: %s", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "challenge not found"})
		return
	}
	log.Printf("Returning challenge with ID: %s", id)
	c.JSON(http.StatusOK, gin.H{"challenge": challenge})
}

func GetChallenges(c *gin.Context) {
	page := 1
	pageSize := 10
	if p := c.Query("page"); p != "" {
		if _, err := fmt.Sscanf(p, "%d", &page); err != nil {
			log.Printf("Invalid page param: %v", err)
			page = 1
		}
	}
	if ps := c.Query("pageSize"); ps != "" {
		if _, err := fmt.Sscanf(ps, "%d", &pageSize); err != nil {
			log.Printf("Invalid pageSize param: %v", err)
			pageSize = 10
		}
	}
	log.Printf("Fetching challenges - page: %d, pageSize: %d", page, pageSize)
	challenges, totalPages, err := repository.GetChallengesWithPagination(c.Request.Context(), page, pageSize)
	if err != nil {
		log.Printf("Error fetching challenges: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch challenges"})
		return
	}
	response := gin.H{
		"currentPage": page,
		"totalPages":  totalPages,
		"challenges":  challenges,
	}
	log.Printf("Response body: %+v", response)
	c.JSON(http.StatusOK, response)
}

func TestChallenge(c *gin.Context) {
	var req types.TestChallengeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind TestChallengeRequest: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if req.ID == 0 || req.LangID == 0 || req.Code == "" || req.SampleInput == "" || req.SampleOutput == "" {
		log.Printf("Validation failed for request: %+v", req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing required fields",
			"details": "All fields (ID, LangID, Code, SampleInput, SampleOutput) must be provided",
		})
		return
	}

	log.Printf("Received TestChallengeRequest: %+v", req)

	c.JSON(http.StatusOK, gin.H{
		"id":            req.ID,
		"lang_id":       req.LangID,
		"code":          req.Code,
		"sample_input":  req.SampleInput,
		"sample_output": req.SampleOutput,
	})
}
