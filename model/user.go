package model

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Password string
}
