package calculator

import (
	"fmt"
	"github.com/alexisleon/stori/internal/models"
	"github.com/alexisleon/stori/internal/util/money"
	"sort"
	"time"
)

// temporary struct used to calculate the TransactionSummaryMonth
type monthCalculation struct {
	date time.Time

	totalCredit,
	totalDebit money.Money

	creditCount,
	debitCount int
}

type ISummaryCalculator interface {
	CalculateSummaryTransaction(r *models.CSVSettlementReport) *models.TransactionSummary
}

var _ ISummaryCalculator = (*SummaryCalculator)(nil)

type SummaryCalculator struct {
	// config, etc
}

func NewSummaryCalculator() *SummaryCalculator {
	return &SummaryCalculator{}
}

func (c SummaryCalculator) CalculateSummaryTransaction(r *models.CSVSettlementReport) *models.TransactionSummary {
	summary := &models.TransactionSummary{
		User: r.User,
	}

	// sort transactions by date
	sort.Slice(r.Transactions, func(i, j int) bool {
		return r.Transactions[i].Date.Before(r.Transactions[j].Date)
	})

	monthsCalc := make(map[string]*monthCalculation)
	for _, txn := range r.Transactions {
		monthID := fmt.Sprintf("%s-%d", txn.Date.Month(), txn.Date.Year())

		if _, exists := monthsCalc[monthID]; !exists {
			monthsCalc[monthID] = &monthCalculation{date: txn.Date}
		}

		if txn.Direction == models.CSVSettlementTransactionDirectionInbound {
			monthsCalc[monthID].totalCredit += txn.Amount()
			monthsCalc[monthID].creditCount += 1
		} else if txn.Direction == models.CSVSettlementTransactionDirectionOutbound {
			monthsCalc[monthID].totalDebit += txn.Amount()
			monthsCalc[monthID].debitCount += 1
		}
	}

	for _, txn := range monthsCalc {
		s := &models.TransactionSummaryMonth{
			Year:              txn.date.Year(),
			Month:             txn.date.Month(),
			TotalTransactions: txn.creditCount + txn.debitCount,
		}
		if txn.creditCount > 0 {
			s.AverageCredit = money.Money(txn.totalCredit.Int() / txn.creditCount)
		}
		if txn.debitCount > 0 {
			s.AverageDebit = money.Money(txn.totalDebit.Int() / txn.debitCount)
		}

		summary.Months = append(summary.Months, s)
		summary.TotalTransactions += txn.creditCount + txn.debitCount
		summary.Balance += txn.totalCredit - txn.totalDebit
	}

	return summary
}
