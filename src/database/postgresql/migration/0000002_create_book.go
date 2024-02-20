package migration

var createBookTableMigration = &Migration{
	Number: 2,
	Name:   "create book table",
	Function: func(code string) error {
		const column = `
			id SERIAL PRIMARY KEY,
			name TEXT
		`
		if code == "update" {
			return Update("book", column)
		}
		if code == "create" {
			return Create("book", column)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createBookTableMigration)
}
