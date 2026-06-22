package main

import (
	"fmt"
	"time"

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


	t := time.NewTicker(1 * time.Hour)

	defer t.Stop()
	done := make(chan bool)

	go func(){
		time.Sleep(12 * time.Hour)
		done <- true
	}()

	for {
		select {
		case <- t.C:
			data := api.GetReport()
			report := excel.ExportToExcel(data)
			if report != nil {
				fmt.Println("Failed to generate report.")
			} else {
				fmt.Println("Report generated successfully.")
			}
		case <- done:
			fmt.Println("Done for the day!")
			return

		}
	}

}