package models

import (
	"github.com/alexisleon/stori/internal/util/money"
	"time"
)

type TransactionSummaryMonth struct {
	Year  int
	Month time.Month

	AverageCredit,
	AverageDebit money.Money

	// The number of transactions during the period
	TotalTransactions int
}

type TransactionSummary struct {
	User              *User
	Account           *Account
	Balance           money.Money
	TotalTransactions int

	Months []*TransactionSummaryMonth
}
