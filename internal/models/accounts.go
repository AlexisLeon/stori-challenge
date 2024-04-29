package models

import (
	"github.com/alexisleon/stori/internal/storage"
	"github.com/alexisleon/stori/internal/util/money"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Account struct {
	ID       uuid.UUID      `db:"id"`
	UserID   string         `db:"user_id"`
	Currency money.Currency `db:"currency"`
	Balance  money.Money    `db:"balance"`
}

func (Account) TableName() string {
	return "accounts"
}

func GetAccount(conn *storage.Conn) (*Account, error) {
	m := &Account{}
	err := conn.Find(m, uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000"))

	if err != nil {
		return nil, errors.Wrap(err, "find account")
	}

	return m, nil
}
