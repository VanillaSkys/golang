package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	RequestId string
	DB        *pgxpool.Pool
}

func New(requestId string, db *pgxpool.Pool) Repository {
	return Repository{RequestId: requestId, DB: db}
}
