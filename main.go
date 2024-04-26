package main

import (
	"encoding/csv"
	"fmt"
	"github.com/alexisleon/stori/internal/calculator"
	"github.com/alexisleon/stori/internal/conf"
	"github.com/alexisleon/stori/internal/models"
	"io"
	"strconv"
	"strings"
	"time"
)

const csvFileContents = `ID,Date,Amount
0,2024-01-01,+60.5
0,2024-01-02,-10.3
4,2024-01-03,-20.46
4,2024-01-04,+10`

func init() {
	conf.LoadConfig()
}

func main() {
	fmt.Println("using default config", conf.GetConfig())

	report := models.CSVSettlementReport{
		User: &models.User{
			ID:    0,
			Email: "test@storicard.com",
		},
		Transactions: make([]*models.CSVSettlementTransaction, 0),
	}

	strInput := strings.NewReader(csvFileContents)
	reader := csv.NewReader(strInput)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			fmt.Println("done")
			break
		}

		if err != nil {
			fmt.Println("failed", err)
			break
		}

		// skip header
		if row[0] == "ID" {
			continue
		}

		date, err := time.Parse("2006-01-02", row[1])
		if err != nil {

			fmt.Println(err)
		}

		amount, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Println(err)
		}

		direction := models.CSVSettlementTransactionDirectionOutbound
		if amount >= 0 {
			direction = models.CSVSettlementTransactionDirectionInbound
		}

		data := models.CSVSettlementTransaction{
			ID:        row[0],
			Date:      date,
			RawAmount: amount,
			Direction: direction,
		}

		report.Transactions = append(report.Transactions, &data)

		// TODO: Save to persistent storage
	}

	summary := calculator.NewSummaryCalculator().CalculateSummaryTransaction(&report)
	// print summary
	fmt.Println("Balance", summary.Balance)
	fmt.Println("TotalTransactions", summary.TotalTransactions)
	for _, month := range summary.Months {
		fmt.Print("\n\n")
		fmt.Println("Date:", month.Year, month.Month)
		fmt.Println("TotalTransactions", month.TotalTransactions)
		fmt.Println("AverageDebit", month.AverageDebit)
		fmt.Println("AverageCredit", month.AverageCredit)
	}
	// send email
}
