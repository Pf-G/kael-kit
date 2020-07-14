package share

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func TablePrint(datas [][]string, headers []string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, v := range datas {
		table.Append(v)
	}
	table.Render() // Send output
}

func TablePrintDemo() {
	datas := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}
	headers := []string{"name", "sign", "rank"}
	TablePrint(datas, headers)
}