package printer

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type TablePrinter struct {
	header []string
	data   [][]string
}

// TODO: use make to set the size of the slice
func NewTablePrinter() *TablePrinter {
	return &TablePrinter{}
}

func (p *TablePrinter) AppendRow(row []string) {
	p.data = append(p.data, row)
}

func (p *TablePrinter) SetHeader(header []string) {
	p.header = header
}

func (p *TablePrinter) Print(writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetAutoFormatHeaders(false)

	table.SetHeader(p.header)
	table.AppendBulk(p.data)
	table.Render()
}
