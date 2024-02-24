package test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}
	defer db.Close()

	// Now you can perform your tests using the database connection
}

// var DBTest *pgxpool.Pool

// func TestDatabaseConnection(t *testing.T) {

// 	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		t.Fatalf("failed to config database connection: %v", err)
// 	}

// 	db, err := pgxpool.ConnectConfig(context.Background(), config)
// 	if err != nil {
// 		t.Fatalf("failed to open database connection: %v", err)
// 	}

// 	// Check the connection
// 	// if err := DB.Ping(context.Background()); err != nil {
// 	// 	DB.Close()
// 	// 	return nil, err
// 	// }
// 	DBTest = db
// }
