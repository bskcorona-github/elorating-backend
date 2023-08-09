package models

type Player struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	EloRating   int    `gorm:"not null" json:"elo_rating"`
	TotalGames  int    `gorm:"not null" json:"total_games"`
	TotalWins   int    `gorm:"not null" json:"total_wins"`
	TotalLosses int    `gorm:"not null" json:"total_losses"`
}

// TeamFormationResult はチーム分け結果を表す構造体です。
type TeamFormationResult struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	TeamA string `json:"team_a" gorm:"-:all"` // プレイヤーのIDのスライスを文字列で保存
	TeamB string `json:"team_b" gorm:"-:all"` // プレイヤーのIDのスライスを文字列で保存
}

// GameResult はゲーム結果を表す構造体です。
type GameResult struct {
	WinningTeam   []*Player
	LosingTeam    []*Player
	WinningPoints int
	LosingPoints  int
}
