package repository

import (
	"context"
	"encoding/json"
	"log"
	"math"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/db"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/types"
)

func GetChallengeByID(ctx context.Context, id string) (*types.Challenge, error) {
	db.Init()

	result, _, err := db.Client.
		From("challenges").
		Select("*", "", false).
		Eq("id", id).
		Single().
		Execute()

	if err != nil {
		log.Printf("Error fetching challenge ID %s: %v", id, err)
		return nil, err
	}

	if result == nil {
		log.Printf("Challenge not found with ID: %s", id)
		return nil, nil
	}

	var challenge types.Challenge
	if err := json.Unmarshal(result, &challenge); err != nil {
		log.Printf("Error unmarshalling challenge ID %s: %v", id, err)
		return nil, err
	}

	log.Printf("Successfully fetched challenge ID: %s", id)
	return &challenge, nil
}

func GetChallengesWithPagination(ctx context.Context, page, pageSize int) ([]types.Challenge, int, error) {
	db.Init()

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	log.Printf("Calculating offset: %d", offset)

	_, count, err := db.Client.
		From("challenges").
		Select("id", "exact", false).
		Execute()

	if err != nil {
		log.Printf("Error counting challenges: %v", err)
		return nil, 0, err
	}

	totalCount := int(count)
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if totalPages == 0 {
		totalPages = 1
	}

	log.Printf("Total challenges: %d, totalPages: %d", totalCount, totalPages)
	result, _, err := db.Client.
		From("challenges").
		Select("*", "", false).
		Range(offset, offset+pageSize-1, "").
		Execute()

	if err != nil {
		log.Printf("Error fetching challenges from DB: %v", err)
		return nil, totalPages, err
	}

	var challenges []types.Challenge
	if err := json.Unmarshal(result, &challenges); err != nil {
		log.Printf("Error unmarshalling challenges: %v", err)
		return nil, totalPages, err
	}

	log.Printf("Returning %d challenges", len(challenges))
	return challenges, totalPages, nil
}
