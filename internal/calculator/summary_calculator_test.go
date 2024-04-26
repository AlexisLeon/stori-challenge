package calculator

import (
	"github.com/alexisleon/stori/internal/models"
	"github.com/alexisleon/stori/internal/util/money"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// I'll keep things simple for this MVP
// But I prefer write tests with Ginkgo + gomega, testify or something similar

func TestNewTransactionSummary(t *testing.T) {
	// transactions get canceled out
	report := &models.CSVSettlementReport{
		Transactions: []*models.CSVSettlementTransaction{
			// January
			{RawAmount: 134.56, Direction: models.CSVSettlementTransactionDirectionInbound, Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 34.56, Direction: models.CSVSettlementTransactionDirectionOutbound, Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 12.34, Direction: models.CSVSettlementTransactionDirectionInbound, Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 12.34, Direction: models.CSVSettlementTransactionDirectionOutbound, Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			// balance=100

			// February
			{RawAmount: 145.67, Direction: models.CSVSettlementTransactionDirectionInbound, Date: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 45.67, Direction: models.CSVSettlementTransactionDirectionOutbound, Date: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 67.89, Direction: models.CSVSettlementTransactionDirectionInbound, Date: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)},
			{RawAmount: 67.89, Direction: models.CSVSettlementTransactionDirectionOutbound, Date: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)},
			// balance=prev+100=200

			// March
			{RawAmount: 100, Direction: models.CSVSettlementTransactionDirectionOutbound, Date: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)},
			// balance=prev-100=100
		},
	}

	calc := NewSummaryCalculator()
	summary := calc.CalculateSummaryTransaction(report)

	assert.Equal(t, money.Money(100_00), summary.Balance)
	assert.Equal(t, 3, len(summary.Months))
	assert.Equal(t, len(report.Transactions), summary.TotalTransactions)

	for _, monthSummary := range summary.Months {
		switch monthSummary.Month {
		case time.January:
			assert.EqualValues(t, 73_45, monthSummary.AverageCredit)
			assert.EqualValues(t, 23_45, monthSummary.AverageDebit)
			assert.Equal(t, 4, monthSummary.TotalTransactions)
		case time.February:
			assert.EqualValues(t, 106_78, monthSummary.AverageCredit)
			assert.EqualValues(t, 56_78, monthSummary.AverageDebit)
			assert.Equal(t, 4, monthSummary.TotalTransactions)
		case time.March:
			assert.EqualValues(t, 0, monthSummary.AverageCredit)
			assert.EqualValues(t, 100_00, monthSummary.AverageDebit)
			assert.Equal(t, 1, monthSummary.TotalTransactions)
		default:
			assert.Fail(t, "unhandled month")
		}
	}
}
