package main

import (
	"encoding/csv"
	"fmt"
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
			Amount:    amount,
			Direction: direction,
		}

		fmt.Println(row, data)

		// TODO: Save to persistent storage
	}
}
