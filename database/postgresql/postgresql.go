package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "go_test"
)

var DB *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	DB = db
	return DB, nil
}
