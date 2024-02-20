package migration

import (
	"fmt"
	"sort"

	"github.com/VanillaSkys/golang/database/postgresql"
)

type Migration struct {
	Number   uint
	Name     string
	Function func(code string) error
}

var Migrations []*Migration

func Migrate(code string) {

	// Sort migrations by number
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Number < Migrations[j].Number
	})

	// Execute migrations
	for _, m := range Migrations {
		m.Function(code)
		// if err := m.Function(code); err != nil {
		// 	fmt.Printf("Error running migration '%s': %s\n", m.Name, err)
		// 	return
		// }
		// fmt.Printf("Migration '%s' executed successfully\n", m.Name)
	}
}

func Create(name string, column string) error {
	db, err := postgresql.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Drop table
	dropSql := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", name)
	if err != nil {
		fmt.Println("Error dropping table:", err)
		return err
	}
	fmt.Printf("Table '%s' dropped successfully!", name)
	// Create table
	db.Exec(dropSql)
	create := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", name, column)
	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return err
	}
	fmt.Printf("Table '%s' created successfully!", name)
	return nil
}

func Update(name string, column string) error {
	db, err := postgresql.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create table
	create := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", name, column)
	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return err
	}
	fmt.Printf("Table '%s' created successfully!", name)
	return nil
}
