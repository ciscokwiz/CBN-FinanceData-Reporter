package api

import (
	"encoding/json"
	"io"
	"net/http"

)


func GetReport() []FinancialData {

	url := "https://www.cbn.gov.ng/api/GetAllFinancialData"


	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var finData []FinancialData

	err = json.Unmarshal(body, &finData)
	if err != nil {
		panic(err)
	}

	return finData
}