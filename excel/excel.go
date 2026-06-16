package excel

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/ciscokwiz/CBN-FinanceData-Reporter/api"

	"github.com/xuri/excelize/v2"
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

	for i := 0; i < value.NumField(); i++ {
		headers[i] = value.Field(i).Name
	}

	headers = append(headers, "Daily_Liquidity")

	return sw.SetRow("A1", headers)
}

func writeRows(sw *excelize.StreamWriter, data []api.FinancialData) error {

	for index, record := range data {

		row := StructToSlice(record)
		dailyLiquidity := calculateDailyLiquidity(record)
		row = append(row, dailyLiquidity)

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

func calculateDailyLiquidity(data api.FinancialData) float64 {

	openbal, _ := strconv.ParseFloat(data.OpeBal, 64)
	slFacility, _ := strconv.ParseFloat(data.SlFacility, 64)

	dailyliquid := openbal * slFacility
	return dailyliquid
}
