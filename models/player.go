package models

type Player struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	EloRating   int    `gorm:"not null" json:"elo_rating"`
	TotalGames  int    `gorm:"not null" json:"total_games"`
	TotalWins   int    `gorm:"not null" json:"total_wins"`
	TotalLosses int    `gorm:"not null" json:"total_losses"`
}
