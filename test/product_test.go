package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/VanillaSkys/golang/database/postgresql/migration"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSave(t *testing.T) {
	migration.Migrate("create")
	// Configure the connection pool
	connPool, err := pgxpool.Connect(context.Background(), "postgresql://username:password@localhost:5432/testdb")
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	defer connPool.Close()

	// Example query
	rows, err := connPool.Query(context.Background(), "SELECT id, name FROM product")
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			t.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("Error iterating through result set: %v", err)
	}

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
