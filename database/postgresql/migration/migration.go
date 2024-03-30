// migration.go
package migration

import (
	"fmt"
	"sort"

	"github.com/VanillaSkys/golang/database/postgresql"
	"gorm.io/gorm"
)

type Migration struct {
	Number   uint
	Name     string
	Function func(db *gorm.DB) error
}

var Migrations []*Migration

// Migrate performs database migrations.
func Migrate(db *gorm.DB, action string) error {
	sortMigrations()

	for _, m := range Migrations {
		if action == "update" && m.Function != nil {
			if err := m.Function(db); err != nil {
				return fmt.Errorf("error running migration '%s': %w", m.Name, err)
			}
			fmt.Printf("Migration '%s' executed successfully\n", m.Name)
		} else if action == "create" && m.Function != nil {
			if err := m.Function(db); err != nil {
				return fmt.Errorf("error running migration '%s': %w", m.Name, err)
			}
			fmt.Printf("Migration '%s' executed successfully\n", m.Name)
		} else {
			return fmt.Errorf("invalid action '%s': must be 'update' or 'create'", action)
		}
	}

	return nil
}

func Create(name string, column string) error {
	type Model struct {
		gorm.Model
	}

	if err := postgresql.DB.AutoMigrate(&Model{}); err != nil {
		return fmt.Errorf("error creating table '%s': %w", name, err)
	}
	fmt.Printf("Table '%s' created successfully\n", name)
	return nil
}

func Update(name string, column string) error {
	type Model struct {
		gorm.Model
	}

	if err := postgresql.DB.AutoMigrate(&Model{}); err != nil {
		return fmt.Errorf("error updating table '%s': %w", name, err)
	}
	fmt.Printf("Table '%s' updated successfully\n", name)
	return nil
}

func sortMigrations() {
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Number < Migrations[j].Number
	})
}
