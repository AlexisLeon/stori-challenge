package models

import (
	"github.com/alexisleon/stori/internal/util/money"
	"math"
	"time"
)

type CSVSettlementTransactionDirection string

const (
	// CSVSettlementTransactionDirectionInbound represents money added to the account
	CSVSettlementTransactionDirectionInbound CSVSettlementTransactionDirection = "inbound"

	// CSVSettlementTransactionDirectionOutbound represents money removed the account
	CSVSettlementTransactionDirectionOutbound CSVSettlementTransactionDirection = "outbound"
)

type CSVSettlementTransaction struct {
	ID   string
	Date time.Time

	RawAmount float64

	Direction CSVSettlementTransactionDirection
}

func (t CSVSettlementTransaction) Amount() money.Money {
	// convert +- to Money
	abs := math.Abs(t.RawAmount)
	return money.Float64(abs)
}

type CSVSettlementReport struct {
	User         *User
	Transactions []*CSVSettlementTransaction
}
