package model

type Product struct {
	Id   string `gorm:"primaryKey"`
	Name string
}
