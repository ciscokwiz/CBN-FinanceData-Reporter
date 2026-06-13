package excel

import (
	"reflect"
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/ciscokwiz/CBN-FinanceData-Reporter/api"
)


func ExportToExcel(data []api.FinancialData) error {

	f := excelize.NewFile()

	sw, err := f.NewStreamWriter("Sheet1")

	if err != nil {
		return err
	}


	err = writeHeaders(sw, data[0])

	if err != nil {
		return err
	}


	err = writeRows(sw, data)

	if err != nil {
		return err
	}


	err = sw.Flush()

	if err != nil {
		return err
	}


	return f.SaveAs("Financial_Data.xlsx")
}


func writeHeaders(sw *excelize.StreamWriter, data api.FinancialData) error {

	value := reflect.TypeOf(data)

	headers := make([]interface{}, value.NumField())

	for i := 0; i < value.NumField(); i++{
		headers[i] = value.Field(i).Name
	}

	return sw.SetRow("A1", headers)
}

func writeRows(sw *excelize.StreamWriter, data []api.FinancialData) error {

	for index, record := range data {

		row := StructToSlice(record)

		cell := fmt.Sprintf("A%d", index+2)

		err := sw.SetRow(cell, row)

		if err != nil {
			return err
		}
	}
	return nil
}

func StructToSlice(data api.FinancialData) []interface{} {

	value := reflect.ValueOf(data)

	row := make(
		[]interface{},
		value.NumField(),
	)

	for i := 0; i < value.NumField(); i++ {

		row[i] = value.Field(i).Interface()
	
	}

	return row
}