// elo_calculator.go

package utils

import (
	"math"
)

// KFactor is the constant factor used in the Elo rating system to determine the influence of each game on the player's rating.
const KFactor = 32

// CalculateElo calculates the new Elo rating for a player based on the results of a game.
func CalculateElo(playerRating, opponentRating, score float64) float64 {
	expectedScore := calculateExpectedScore(playerRating, opponentRating)
	newRating := playerRating + KFactor*(score-expectedScore)
	return math.Round(newRating)
}

// calculateExpectedScore calculates the expected score of a player against an opponent based on their ratings.
func calculateExpectedScore(playerRating, opponentRating float64) float64 {
	return 1.0 / (1.0 + math.Pow(10, (opponentRating-playerRating)/400))
}
