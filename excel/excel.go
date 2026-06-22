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

	headers = append(headers, "Liquidity (N'000,000)")

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
	rediscbills,_ := strconv.ParseFloat(data.RediscBills, 64)
	slfacility, _ := strconv.ParseFloat(data.SlFacility, 64)
	sdfacility, _ := strconv.ParseFloat(data.SdFacility, 64)
	repo, _ := strconv.ParseFloat(data.Repo, 64)
	revrepo,_ := strconv.ParseFloat(data.RevRepo, 64)
	omosales, _ := strconv.ParseFloat(data.OmoSales, 64)
	omorepay,_ := strconv.ParseFloat(data.OmoRepay, 64)
	pmsales,_ := strconv.ParseFloat(data.PmSales, 64)
	pmrepay,_ := strconv.ParseFloat(data.PmRepay, 64) 
	crr, _ := strconv.ParseFloat(data.Crr, 64)
	netwdas, _ := strconv.ParseFloat(data.NetWdas, 64)
	statalloc,_ := strconv.ParseFloat(data.StatAlloc, 64)
	jvcash, _ := strconv.ParseFloat(data.JvCash, 64)
	netclr,_ := strconv.ParseFloat(data.NetClr, 64)
	ndicprem, _ := strconv.ParseFloat(data.NdicPrem, 64)
	omajor, _ := strconv.ParseFloat(data.OMajor, 64)



	Liquidity := (openbal + rediscbills - slfacility + sdfacility - repo + revrepo - omosales + omorepay - pmsales + pmrepay - crr + netwdas - statalloc + jvcash - netclr + ndicprem - omajor) * 1000000


	return Liquidity
}
