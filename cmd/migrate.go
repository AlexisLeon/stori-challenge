package cmd

import (
	"github.com/alexisleon/stori/internal/conf"
	"github.com/alexisleon/stori/internal/storage"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
)

var migrateCmd = cobra.Command{
	Use:  "migrate",
	Long: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

// I would normally use something like github.com/golang-migrate/migrate for migrations
// In this case I wanted to try out https://github.com/gobuffalo/pop
func migrate() {
	c := conf.LoadConfig(configFile)

	db, err := storage.Connect(c)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "failed to open db connection"))
	}
	defer db.Close()

	if err := db.Open(); err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "checking database connection"))
	}

	migrationsPath := "db/migrations"
	log.Printf("Reading migrations from %s\n", migrationsPath)
	fileMigrator, err := pop.NewFileMigrator(migrationsPath, db.Connection)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "failed to migrate db"))
	}

	err = fileMigrator.Status()
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "migration status"))
	}

	// Migrate
	err = fileMigrator.Up()
	if err != nil {
		log.Fatalf("%v", errors.Wrap(err, "running db migrations"))
	} else {
		log.Println("migrations applied successfully")
	}

	err = fileMigrator.Status()
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "migration status"))
	}
}
