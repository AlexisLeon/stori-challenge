package csv_reader

import (
	"fmt"
	"github.com/alexisleon/stori/internal/models"
	"github.com/gofrs/uuid"
	"strconv"
	"time"
)

type settlementReportReader struct {
	Transactions []*models.CSVSettlementTransaction
}

type SettlementReportReader = Reader[*models.CSVSettlementTransaction]

var _ Reader[*models.CSVSettlementTransaction] = (*settlementReportReader)(nil)

func NewSettlementReportReader() Reader[*models.CSVSettlementTransaction] {
	return &settlementReportReader{
		Transactions: []*models.CSVSettlementTransaction{},
	}
}

func (r *settlementReportReader) ReadRow(columns []string) error {
	if len(columns) != 3 {
		return fmt.Errorf("invalid row, expected 3 columns, got %d", len(columns))
	}

	// skip header
	if columns[0] == "ID" {
		return nil
	}

	date, err := time.Parse("2006-01-02", columns[1])
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(columns[2], 64)
	if err != nil {
		return err
	}

	direction := models.CSVSettlementTransactionDirectionOutbound
	if amount >= 0 {
		direction = models.CSVSettlementTransactionDirectionInbound
	}

	id, err := uuid.FromString(columns[0])
	if err != nil {
		return err
	}

	data := models.CSVSettlementTransaction{
		ID:        id,
		Date:      date,
		RawAmount: amount,
		Direction: direction,
	}

	r.Transactions = append(r.Transactions, &data)

	return nil
}

func (r settlementReportReader) Rows() []*models.CSVSettlementTransaction {
	return r.Transactions
}
