package service

import (
	"encoding/csv"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type CurrencyPair struct {
	FromCurrency string
	ToCurrency   string
}

type CurrencyIngestor struct {
	currencyManager ICurrencyManager
	cfg             *Config
}

func NewCurrencyIngestor(
	currencyManager ICurrencyManager,
	cfg *Config,
) *CurrencyIngestor {
	return &CurrencyIngestor{
		currencyManager: currencyManager,
		cfg:             cfg,
	}
}

func (ingestor *CurrencyIngestor) Truncate() error {
	err := ingestor.currencyManager.Truncate()
	if err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	return nil
}

func (ingestor *CurrencyIngestor) ProcessCurrencyRates() error {
	log.Println("Processing currency rates...")

	log.Printf("%+v\n", ingestor.cfg)
	dirPath := ingestor.cfg.CurrencyRate.DirPath

	// Read the directory contents.
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate through the files and print their names.
	for _, file := range files {
		fmt.Println(file.Name())
		filePath := dirPath + "/" + file.Name()

		pair, err := getCurrencyPair(file.Name())
		if err != nil {
			return err
		}

		err = ingestor.processCurrencyRates(filePath, pair)
		if err != nil {
			return err
		}
	}

	return nil
}

// getCurrencyPair extracts the currency pair from a given file name based on its format and separator.
// e.g. USD.AUD.csv.
func getCurrencyPair(fileName string) (*CurrencyPair, error) {
	// Remove the file extension.
	fileNameWithoutExt := strings.TrimSuffix(fileName, ".csv")

	// Split the string by the underscore.
	parts := strings.Split(fileNameWithoutExt, ".")

	// Check if the file name follows the expected format.
	if len(parts) != 2 {
		return nil, fmt.Errorf("cannot parse currency pair from file name: %s", fileName)
	}

	return &CurrencyPair{
		FromCurrency: parts[0],
		ToCurrency:   parts[1],
	}, nil
}

func (ingestor *CurrencyIngestor) processCurrencyRates(filePath string, pair *CurrencyPair) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file: %s\n", err.Error())
	}

	defer file.Close()

	// Create a new CSV reader.
	reader := csv.NewReader(file)

	// Read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading CSV data: %s\n", err.Error())
	}

	// Skip the header row (first row).
	if len(records) > 0 {
		records = records[1:]
	}

	dateFormat := "2006-01-02"

	for _, row := range records {
		// Parse the date string into a time.Time object.
		tradeDate, err := time.Parse(dateFormat, row[0])
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			continue
		}

		closePrice, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Printf("Error parsing Close value: %v\n", err)
			continue
		}

		currencyRate := model.CurrencyRate{
			BaseCurrency:   pair.FromCurrency,
			TargetCurrency: pair.ToCurrency,
			ExchangeRate:   closePrice,
			TradeDate:      &tradeDate,
		}

		err = ingestor.currencyManager.AddCurrencyRate(&currencyRate)
		if err != nil {
			return err
		}

	}

	return nil
}
