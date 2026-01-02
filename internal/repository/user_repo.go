package repository

import (
	"context"
	"encoding/json"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/db"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/types"
)

func GetUserByID(ctx context.Context, id string) (*types.User, error) {
	db.Init()

	result, _, err := db.Client.
		From("users").
		Select("*", "", false).
		Eq("id", id).
		Single().
		Execute()

	if err != nil {
		return nil, err
	}

	var user types.User
	if err := json.Unmarshal(result, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
