package cmd

import (
	"bytes"
	"github.com/alexisleon/stori/internal/calculator"
	"github.com/alexisleon/stori/internal/conf"
	"github.com/alexisleon/stori/internal/csv_reader"
	"github.com/alexisleon/stori/internal/mailer/templates"
	"github.com/alexisleon/stori/internal/models"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var templateCmd = cobra.Command{
	Use: "template",
	Run: func(cmd *cobra.Command, args []string) {
		template()
	},
}

func template() {
	conf.LoadConfig(configFile)
	settlementReader := csv_reader.NewSettlementReportReader()
	err := csv_reader.ReadCSVFile("settlement.csv", settlementReader)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	report := models.CSVSettlementReport{
		User:         &models.User{Email: "alexisleon@storicard.com"},
		Transactions: settlementReader.Rows(),
	}

	summary := calculator.NewSummaryCalculator().CalculateSummaryTransaction(&report)
	summary.Account = &models.Account{
		Currency: "MXN",
		Balance:  12345,
	}

	msg := new(bytes.Buffer)
	tpl := templates.NewSummaryMailerTemplate()
	err = tpl.Compile(msg, summary)
	if err != nil {
		log.Fatalf("failed to compile template: %v", err)
	}

	err = os.WriteFile("summary.html", msg.Bytes(), 0600)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}

	log.Println("done")
}
