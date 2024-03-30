// cmd.go
package cmd

import (
	"fmt"
	"os"

	"github.com/VanillaSkys/golang/database/postgresql"
	"github.com/VanillaSkys/golang/database/postgresql/migration"
	"github.com/VanillaSkys/golang/interface/http"
)

// Execute is the main entry point for command execution.
func Execute() {
	// Connect to the database
	postgresql.ConnectDB()

	// If no command-line arguments provided, start HTTP server or print usage
	if len(os.Args) == 1 {
		fmt.Println("No command specified. Usage: cmd [migration [update|create]]")
		http.InitHttpServer()
		os.Exit(1)
	}

	// Check for "migration" command
	if os.Args[1] == "migration" && len(os.Args) >= 3 {
		// Check for "update" or "create" sub-command
		if os.Args[2] == "update" || os.Args[2] == "create" {
			// Perform migration
			err := migration.Migrate(postgresql.DB, os.Args[2])
			if err != nil {
				fmt.Printf("Error performing migration: %v\n", err)
				os.Exit(1)
			}
			return
		}
	}

	// Print usage if command-line arguments are not recognized
	fmt.Println("Usage: cmd [migration [update|create]]")
	os.Exit(1)
}
