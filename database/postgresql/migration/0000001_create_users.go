package migration

var createUserTableMigration = &Migration{
	Number: 1,
	Name:   "create users table",
	Function: func(code string) error {
		const column = `
			id TEXT PRIMARY KEY,
			username TEXT
		`
		if code == "update" {
			return Update("users", column)
		}
		if code == "create" {
			return Create("users", column)
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, createUserTableMigration)
}
