package cmd

import (
	"os"

	"github.com/VanillaSkys/golang/database/postgresql"
	"github.com/VanillaSkys/golang/database/postgresql/migration"
	"github.com/VanillaSkys/golang/interface/http"
)

func Execute() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	postgresql.ConnectDB()
	if len(os.Args) == 1 {
		http.InitHttpServer()
	} else {
		if os.Args[1] == "migration" {

			if os.Args[2] == "update" {
				migration.Migrate("update")
			}
			if os.Args[2] == "create" {
				migration.Migrate("create")
			}
		}
	}
}
