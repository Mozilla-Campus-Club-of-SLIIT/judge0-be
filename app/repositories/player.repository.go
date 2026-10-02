package repositories

import (
	"context"
	"errors"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/database"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/logger"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/types"
	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/utils"
)

var ErrPlayerNotFound = errors.New("player not found")

// GetAllPlayers returns every registered user with their current marks,
// intended for a frontend player picker (no pagination needed for this use case).
func GetAllPlayers(ctx context.Context) ([]types.PlayerType, error) {
	pool := database.GetPool()
	ctx, cancel := utils.WithTimeout(ctx)
	defer cancel()

	rows, err := pool.Query(ctx,
		`SELECT u.user_id, u.name, u.email, COALESCE(l.marks, 0)
		 FROM users u
		 LEFT JOIN leaderboard l ON l.user_id = u.user_id
		 ORDER BY u.name ASC`,
	)
	if err != nil {
		logger.Log.Error("GetAllPlayers: query error", "error", err)
		return nil, err
	}
	defer rows.Close()

	players := []types.PlayerType{}
	for rows.Next() {
		var player types.PlayerType
		if err := rows.Scan(&player.UserID, &player.Name, &player.Email, &player.Marks); err != nil {
			logger.Log.Error("GetAllPlayers: scan error", "error", err)
			return nil, err
		}
		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Error("GetAllPlayers: rows error", "error", err)
		return nil, err
	}

	logger.Log.Info("Fetched players", "count", len(players))
	return players, nil
}

// AddMarksToPlayer adjusts a player's leaderboard marks by delta (positive or negative)
// and returns the resulting total. It fails with ErrPlayerNotFound if the user doesn't exist.
func AddMarksToPlayer(ctx context.Context, userID string, delta int) (int, error) {
	pool := database.GetPool()
	ctx, cancel := utils.WithTimeout(ctx)
	defer cancel()

	tx, err := pool.Begin(ctx)
	if err != nil {
		logger.Log.Error("AddMarksToPlayer: begin transaction error", "user_id", userID, "error", err)
		return 0, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var userExists bool
	err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userID).Scan(&userExists)
	if err != nil {
		logger.Log.Error("AddMarksToPlayer: user existence check error", "user_id", userID, "error", err)
		return 0, err
	}
	if !userExists {
		logger.Log.Warn("AddMarksToPlayer: player not found", "user_id", userID)
		return 0, ErrPlayerNotFound
	}

	var total int
	err = tx.QueryRow(ctx,
		`INSERT INTO leaderboard (user_id, marks)
		 VALUES ($1, $2)
		 ON CONFLICT (user_id) DO UPDATE SET marks = leaderboard.marks + $2, last_updates = now()
		 RETURNING marks`,
		userID, delta,
	).Scan(&total)
	if err != nil {
		logger.Log.Error("AddMarksToPlayer: upsert error", "user_id", userID, "delta", delta, "error", err)
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Log.Error("AddMarksToPlayer: commit error", "user_id", userID, "error", err)
		return 0, err
	}

	logger.Log.Info("AddMarksToPlayer: marks adjusted", "user_id", userID, "delta", delta, "total", total)
	return total, nil
}
