package service

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func getDateValue(dateString string) (*time.Time, error) {
	if dateString == "" {
		return nil, nil
	}

	dateValue, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}

	return &dateValue, nil
}

func getFloatValue(valueString string) (float64, error) {
	if valueString == "" {
		return 0, nil
	}

	value, err := strconv.ParseFloat(normalizeNumber(valueString), 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func normalizeNumber(s string) string {
	return strings.Replace(s, ",", "", -1)
}

func getStringValue(valueString string) string {
	if valueString == "" {
		return ""
	}

	return strings.ToUpper(strings.TrimSpace(valueString))
}

func displayRow(row []string) {
	log.Printf("row: %v", row)
}
