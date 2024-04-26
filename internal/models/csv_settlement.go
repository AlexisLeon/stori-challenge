package models

import "time"

type CSVSettlementTransactionDirection string

const (
	// CSVSettlementTransactionDirectionInbound represents money added to the account
	CSVSettlementTransactionDirectionInbound CSVSettlementTransactionDirection = "inbound"

	// CSVSettlementTransactionDirectionOutbound represents money removed the account
	CSVSettlementTransactionDirectionOutbound CSVSettlementTransactionDirection = "outbound"
)

type CSVSettlementTransaction struct {
	ID        string
	Date      time.Time
	Amount    float64
	Direction CSVSettlementTransactionDirection
}

type CSVSettlementReport struct {
	User        *User
	Transaction []CSVSettlementTransaction
}
