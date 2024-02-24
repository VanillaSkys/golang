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

	DB, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Check the connection
	// if err := DB.Ping(context.Background()); err != nil {
	// 	DB.Close()
	// 	return nil, err
	// }
	return DB, nil
}
