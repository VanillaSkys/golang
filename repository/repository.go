package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	RequestId string
	DB        *gorm.DB
}

func New(requestId string, db *gorm.DB) Repository {
	return Repository{RequestId: requestId, DB: db}
}
