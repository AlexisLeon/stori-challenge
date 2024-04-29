package storage

import (
	"github.com/alexisleon/stori/internal/conf"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

type Conn struct {
	*pop.Connection
}

func Connect(c *conf.Config) (*Conn, error) {
	connDets := &pop.ConnectionDetails{
		Dialect: "postgres",
		URL:     c.Database.URL,
		// options that will be passed to each migration file
		Options: map[string]string{
			"Namespace":            "alexisleon_stori",
			"migration_table_name": "schema_migrations",
		},

		// Add Pool size, idle, max lifetime
	}

	conn, err := pop.NewConnection(connDets)
	if err != nil {
		return nil, errors.Wrap(err, "failed to establish database connection")
	}

	if err := conn.Open(); err != nil {
		return nil, errors.Wrap(err, "unable to open database connection")
	}

	return &Conn{conn}, nil
}
