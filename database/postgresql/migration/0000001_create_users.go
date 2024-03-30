package migration

import (
	"github.com/VanillaSkys/golang/model"
	"gorm.io/gorm"
)

var createUserTableMigration = &Migration{
	Number: 1,
	Name:   "create or update users table",
	Function: func(db *gorm.DB) error {

		// Check if the "users" table already exists
		if db.Migrator().HasTable(&model.User{}) {
			// Table exists, update it
			if err := db.AutoMigrate(&model.User{}); err != nil {
				return err
			}
			println("Table 'users' updated successfully")
		} else {
			// Table doesn't exist, create it
			if err := db.AutoMigrate(&model.User{}); err != nil {
				return err
			}
			println("Table 'users' created successfully")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createUserTableMigration)
}
