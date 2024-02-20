package migration

var createBookTableMigration = &Migration{
	Number: 2,
	Name:   "create books table",
	Function: func(code string) error {
		const column = `
			id SERIAL PRIMARY KEY,
			name TEXT
		`
		if code == "update" {
			return Update("books", column)
		}
		if code == "create" {
			return Create("books", column)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createBookTableMigration)
}
