package migration

var createProductTableMigration = &Migration{
	Number: 3,
	Name:   "create products table",
	Function: func(code string) error {
		const column = `
			id SERIAL PRIMARY KEY,
			name TEXT
		`
		if code == "update" {
			return Update("products", column)
		}
		if code == "create" {
			return Create("products", column)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createProductTableMigration)
}
