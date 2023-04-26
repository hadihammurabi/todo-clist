package pkg

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func NewTable(columns ...any) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	t := table.New(columns...)
	t.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return t
}
