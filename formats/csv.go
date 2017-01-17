package formats

import (
	"github.com/maxlinc/tabler"
	"io"
	"encoding/csv"
)

type CSV struct {
	Reader io.Reader
	Writer io.Writer
}

func (CSV CSV) Render(table tabler.Table) error {
	csvWriter := csv.NewWriter(CSV.Writer)
	err := csvWriter.Write(table.Headers); if err != nil {
		return err
	}
	csvWriter.WriteAll(table.Rows); if err != nil {
		return err
	}
	return nil
}

func ReadCSV(reader io.Reader) (*tabler.Table, error) {
	var headers []string
	var rows [][]string
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if (err != nil) {
		return nil, err
	}

	for i, row := range records {
		if i == 0 {
			headers = row
		} else {
			rows = append(rows, row)
		}
	}
	table := tabler.Table{
		Headers: headers,
		Rows: rows,
	}
	return &table, nil
}