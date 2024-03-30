package migration

import (
	"github.com/VanillaSkys/golang/model"
	"gorm.io/gorm"
)

var createBookTableMigration = &Migration{
	Number: 2,
	Name:   "create or update books table",
	Function: func(db *gorm.DB) error {

		// Check if the "books" table already exists
		if db.Migrator().HasTable(&model.Book{}) {
			// Table exists, update it
			if err := db.AutoMigrate(&model.Book{}); err != nil {
				return err
			}
			println("Table 'books' updated successfully")
		} else {
			// Table doesn't exist, create it
			if err := db.AutoMigrate(&model.Book{}); err != nil {
				return err
			}
			println("Table 'books' created successfully")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createBookTableMigration)
}
