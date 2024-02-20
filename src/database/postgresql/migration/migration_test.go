package migration

import (
	"testing"
)

func TestCreateUserTableMigration(t *testing.T) {
	const column = `
			id SERIAL PRIMARY KEY,
			username TEXT,
	`
	// Create table

	if err := Create("user", column); err != nil {
		t.Errorf("Migration user error: %v", err)
	}
}
