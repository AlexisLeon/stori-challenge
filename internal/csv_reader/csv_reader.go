package csv_reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Reader[Row any] interface {
	// ReadRow returns a slice of strings representing the columns of the row
	// Implementers should return error if the row is invalid
	ReadRow(columns []string) error

	Rows() []Row
}

func ReadCSVFile[Row any](filePath string, reader Reader[Row]) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(file)

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			fmt.Println("done")
			break
		}

		if err != nil {
			fmt.Println("failed", err)
			break
		}

		err = reader.ReadRow(row)
		if err != nil {
			return fmt.Errorf("failed to read row: %w", err)
		}
	}

	return nil
}
