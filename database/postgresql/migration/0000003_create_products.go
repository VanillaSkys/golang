package migration

import (
	"github.com/VanillaSkys/golang/model"
	"gorm.io/gorm"
)

var createProductTableMigration = &Migration{
	Number: 3,
	Name:   "create or update products table",
	Function: func(db *gorm.DB) error {

		// Check if the "products" table already exists
		if db.Migrator().HasTable(&model.Product{}) {
			// Table exists, update it
			if err := db.AutoMigrate(&model.Product{}); err != nil {
				return err
			}
			println("Table 'products' updated successfully")
		} else {
			// Table doesn't exist, create it
			if err := db.AutoMigrate(&model.Product{}); err != nil {
				return err
			}
			println("Table 'products' created successfully")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createProductTableMigration)
}
