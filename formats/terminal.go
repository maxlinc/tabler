package formats

import (
	"io"
	"github.com/maxlinc/tabler"
	"github.com/olekukonko/tablewriter"
)

type Terminal struct {
	Writer io.Writer
}

func (terminal Terminal) Render(table tabler.Table) error {
	tableWriter := tablewriter.NewWriter(terminal.Writer)
	tableWriter.SetHeader(table.Headers)
	tableWriter.AppendBulk(table.Rows)
	tableWriter.Render()
	return nil
}
