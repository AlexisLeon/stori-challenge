package models

import (
	"github.com/alexisleon/stori/internal/storage"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type User struct {
	ID    uuid.UUID `db:"id"`
	Email string    `db:"email"`
}

func (User) TableName() string {
	return "users"
}

func GetUser(conn *storage.Conn) (*User, error) {
	// for this project, only have one user
	user := &User{}
	err := conn.Find(user, uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000"))

	if err != nil {
		return nil, errors.Wrap(err, "find user")
	}

	return user, nil
}
