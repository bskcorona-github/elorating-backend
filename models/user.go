package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
