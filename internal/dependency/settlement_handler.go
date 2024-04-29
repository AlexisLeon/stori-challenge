package dependency

import (
	"github.com/alexisleon/stori/internal/calculator"
	"github.com/alexisleon/stori/internal/csv_reader"
	"github.com/alexisleon/stori/internal/mailer"
	"github.com/alexisleon/stori/internal/models"
	"github.com/alexisleon/stori/internal/storage"
	"github.com/pkg/errors"
	"log"
)

type SettlementHandler struct {
	db     *storage.Conn
	mailer mailer.Mailer
}

func NewSettlementHandler(c *storage.Conn, m mailer.Mailer) *SettlementHandler {
	return &SettlementHandler{
		db:     c,
		mailer: m,
	}
}

func (h *SettlementHandler) ProcessSettlement() error {
	settlementReader := csv_reader.NewSettlementReportReader()
	err := csv_reader.ReadCSVFile("settlement.csv", settlementReader)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	user, err := models.GetUser(h.db)
	if err != nil {
		return errors.Wrap(err, "get user")
	}

	// start transaction
	for _, ftxn := range settlementReader.Rows() {
		exists, rerr := models.GetTransactionByTransactionUUID(h.db, ftxn.ID)
		if rerr != nil {
			return errors.Wrap(err, "get transaction")
		}

		if exists != nil {
			continue
		}

		rerr = h.db.StartTransaction(func(conn *storage.Conn) error {
			account, err := models.GetAccount(h.db)
			if err != nil {
				return errors.Wrap(err, "get account")
			}

			// You'll normally want to insert the transaction once and make this step idempotent
			// but for this example we'll keep it simple
			txn, err := models.NewTransaction(user.ID, account.ID, ftxn.ID, ftxn.Amount())

			if err != nil {
				return errors.Wrap(err, "new transaction")
			}

			err = conn.Create(txn)
			if err != nil {
				return errors.Wrap(err, "create transaction")
			}

			// You'll normally want to update the balance in a more robust way
			// that guarantee the integrity of the ledger, but for this example,
			// we'll just update the balance by adding the amount

			if ftxn.Direction == models.CSVSettlementTransactionDirectionOutbound {
				account.Balance -= ftxn.Amount()
			} else if ftxn.Direction == models.CSVSettlementTransactionDirectionInbound {
				account.Balance += ftxn.Amount()
			}

			err = conn.UpdateColumns(account, "balance")
			if err != nil {
				return errors.Wrap(err, "update account balance")
			}

			return nil
		})

		if rerr != nil {
			return errors.Wrap(err, "transaction")
		}
	}

	report := models.CSVSettlementReport{
		User:         user,
		Transactions: settlementReader.Rows(),
	}

	summary := calculator.NewSummaryCalculator().CalculateSummaryTransaction(&report)

	err = h.mailer.SendSummary(summary)
	if err != nil {
		return errors.Wrap(err, "send summary")
	}

	return nil
}
