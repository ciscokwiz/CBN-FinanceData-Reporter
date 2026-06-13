package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


func GetReport() []FinancialData {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	url := os.Getenv("CBN_URL")


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