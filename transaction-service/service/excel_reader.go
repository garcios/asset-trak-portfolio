package service

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func getRows(filePath string, tabName string) ([][]string, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Processing file %s...\n", file.Path)

	return file.GetRows(tabName)
}
