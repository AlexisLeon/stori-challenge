package cmd

import (
	"github.com/alexisleon/stori/internal/conf"
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
	globalConf := conf.LoadConfig(configFile)

	connDets := &pop.ConnectionDetails{
		Dialect: "postgres",
		URL:     globalConf.Database.URL,
		// options that will be passed to each migration file
		Options: map[string]string{
			"Namespace":            "alexisleon_stori",
			"migration_table_name": "schema_migrations",
		},
	}

	db, err := pop.NewConnection(connDets)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "failed to open db connection"))
	}
	defer db.Close()

	if err := db.Open(); err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "checking database connection"))
	}

	migrationsPath := "db/migrations"
	log.Printf("Reading migrations from %s\n", migrationsPath)
	fileMigrator, err := pop.NewFileMigrator(migrationsPath, db)
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
