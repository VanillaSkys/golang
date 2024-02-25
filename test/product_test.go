package test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSave(t *testing.T) {

	// Create a new PostgreSQL connection pool
	pool, err := pgxpool.Connect(context.Background(), "DATABASE_URL")
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// Test your database connection here
	// For example, you can perform a query to ensure the connection is working
	rows, err := pool.Query(context.Background(), "SELECT 1")
	if err != nil {
		t.Fatalf("Error querying database: %v\n", err)
	}
	defer rows.Close()

}

// func TestSave_Error(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	pool, err := pgxpool.Connect(context.Background(), "mock_db_connection_string")
// 	if err != nil {
// 		t.Fatalf("failed to connect to pgxpool: %v", err)
// 	}

// 	r := NewRepository(pool)

// 	name := "test product"

// 	mock.ExpectExec("INSERT INTO product").
// 		WithArgs(name).
// 		WillReturnError(errors.New("error saving product"))

// 	if err := r.Save(name); err == nil {
// 		t.Error("expected error, got nil")
// 	} else {
// 		expectedErr := "error saving product"
// 		if err.Error() != expectedErr {
// 			t.Errorf("expected error '%s', got '%s'", expectedErr, err.Error())
// 		}
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %v", err)
// 	}
// }
