package test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VanillaSkys/golang/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pool, err := pgxpool.Connect(context.Background(), "mock_db_connection_string")
	if err != nil {
		t.Fatalf("failed to connect to pgxpool: %v", err)
	}

	r := repository.New("product_test_123", pool)

	name := "test product"

	mock.ExpectExec("INSERT INTO product").
		WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)

	if err := r.Save(name); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
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

// 	r := repository.New("product_test_123", pool)

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
