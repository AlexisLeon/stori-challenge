package dependency

import (
	"github.com/alexisleon/stori/internal/conf"
	"github.com/alexisleon/stori/internal/mailer"
	"github.com/alexisleon/stori/internal/mailer/provider"
	"github.com/alexisleon/stori/internal/storage"
)

type AppContext struct {
	dbConnPool *storage.Conn
	config     *conf.Config

	Mailer mailer.Mailer

	SettlementHandler *SettlementHandler
}

func CreateAppContext(c *conf.Config) (*AppContext, error) {
	conn, err := storage.Connect(c)
	if err != nil {
		return nil, err
	}

	return NewAppContext(c, conn), nil
}

func NewAppContext(c *conf.Config, conn *storage.Conn) *AppContext {
	mail := mailer.NewMailer(provider.NewAwsSESProvider(c))

	return &AppContext{
		dbConnPool:        conn,
		config:            c,
		Mailer:            mail,
		SettlementHandler: NewSettlementHandler(conn, mail),
	}
}
