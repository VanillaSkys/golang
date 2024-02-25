package test

import (
	"context"
	"os"
	"testing"

	"github.com/VanillaSkys/golang/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSave(t *testing.T) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to config database connection: %v", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}

	r := repository.New("123", db)

	id := "test"
	name := "test"

	if err := r.Save(id, name); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

}

// func TestSave(t *testing.T) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	fmt.Println(os.Getenv("DATABASE_URL"))
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		t.Fatalf("failed to connect to pgxpool: %v", err)
// 	}

// 	r := repository.New("123", pool)

// 	id := "test product"
// 	name := "test product"

// 	mock.ExpectExec("INSERT INTO product ").
// 		WithArgs(name).
// 		WillReturnResult(sqlmock.NewResult(1, 1)).
// 		WillReturnError(nil)

// 	if err := r.Save(id, name); err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %v", err)
// 	}
// }

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
