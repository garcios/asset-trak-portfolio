package typesutils

import (
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

func GetDateValue(dateString string, format string) (*time.Time, error) {
	if dateString == "" {
		return nil, nil
	}

	if format == "" {
		format = "2006-01-02"
	}

	dateValue, err := time.Parse(format, dateString)
	if err != nil {
		return nil, err
	}

	return &dateValue, nil
}

func GetFloatAsDate(valueString string) (*time.Time, error) {
	floatValue, err := GetFloatValue(valueString)
	if err != nil {
		return nil, err
	}

	dateValue, err := excelize.ExcelDateToTime(floatValue, false)
	if err != nil {
		return nil, err
	}

	return &dateValue, nil
}

func GetFloatValue(valueString string) (float64, error) {
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

func GetStringValue(valueString string) string {
	if valueString == "" {
		return ""
	}

	return strings.ToUpper(strings.TrimSpace(valueString))
}
