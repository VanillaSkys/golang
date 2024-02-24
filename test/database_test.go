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
