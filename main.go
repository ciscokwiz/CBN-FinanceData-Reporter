package main

import (
	"fmt"

	"github.com/ciscokwiz/CBN-FinanceData-Reporter/api"
	"github.com/ciscokwiz/CBN-FinanceData-Reporter/excel"
)

func main() {
	data := api.GetReport()
	report := excel.ExportToExcel(data)
	if report != nil {
		fmt.Println("Failed to generate report.")
	} else {
		fmt.Println("Report generated successfully.")
	}
}