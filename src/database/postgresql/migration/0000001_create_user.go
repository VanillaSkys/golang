package migration

var createUserTableMigration = &Migration{
	Number: 1,
	Name:   "create user table",
	Function: func(code string) error {
		const column = `
			id SERIAL PRIMARY KEY,
			username TEXT,
			password TEXT
		`
		if code == "update" {
			return Update("user", column)
		}
		if code == "create" {
			return Create("user", column)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createUserTableMigration)
}
