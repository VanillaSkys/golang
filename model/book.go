package model

type Book struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
