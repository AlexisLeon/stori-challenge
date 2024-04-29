package models

import (
	"github.com/alexisleon/stori/internal/storage"
	"github.com/alexisleon/stori/internal/util/money"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Transaction struct {
	ID              uuid.UUID   `db:"id"`
	UserID          uuid.UUID   `db:"user_id"`
	AccountID       uuid.UUID   `db:"account_id"`
	TransactionUUID uuid.UUID   `db:"transaction_uuid"`
	Amount          money.Money `db:"amount"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func NewTransaction(userID uuid.UUID, accountID uuid.UUID, transactionUUID uuid.UUID, amount money.Money) (*Transaction, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "generate uuid")
	}

	return &Transaction{
		ID:              id,
		UserID:          userID,
		AccountID:       accountID,
		TransactionUUID: transactionUUID,
		Amount:          amount,
	}, nil
}

// GetTransactionByTransactionUUID returns a transaction by transaction uuid, otherwise nil
func GetTransactionByTransactionUUID(conn *storage.Conn, transactionUUID uuid.UUID) (*Transaction, error) {
	m := &Transaction{}
	err := conn.Where("transaction_uuid = ?", transactionUUID).First(m)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, errors.Wrap(err, "find transaction")
	}

	return m, nil
}
