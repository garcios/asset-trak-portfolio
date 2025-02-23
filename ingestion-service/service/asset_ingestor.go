package service

import (
	"encoding/csv"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
	"log"
	"os"

	"github.com/google/uuid"
)

type AssetIngestor struct {
	assetManager IAssetManager
	cfg          *Config
}

func NewAssetIngestor(am IAssetManager, cfg *Config) *AssetIngestor {
	return &AssetIngestor{
		assetManager: am,
		cfg:          cfg,
	}
}

func (ingestor *AssetIngestor) Truncate() error {
	err := ingestor.assetManager.Truncate()
	if err != nil {
		return fmt.Errorf("failed to truncate asset data: %w", err)
	}

	log.Println("truncated asset data successfully")

	return nil
}

func (ingestor *AssetIngestor) ProcessAssets() error {
	log.Println("Processing assets...")

	// Open the CSV file
	file, err := os.Open(ingestor.cfg.Asset.Path)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err.Error())
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %s", err.Error())
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < ingestor.cfg.Asset.SkipRows {
			rowCount++
			continue
		}

		if len(row) == 0 {
			break
		}

		asset, err := mapColumnsToAsset(row)
		if err != nil {
			return err
		}

		isFound, err := ingestor.assetManager.AssetExists(asset.Symbol, asset.MarketCode)
		if err != nil {
			return err
		}

		if isFound {
			continue
		}

		err = ingestor.assetManager.AddAsset(asset)
		if err != nil {
			return err
		}

	}

	return nil
}

func mapColumnsToAsset(row []string) (*model.Asset, error) {
	id := uuid.New()
	rec := &model.Asset{ID: id.String(), Symbol: row[0], Name: row[1], MarketCode: row[2]}

	return rec, nil

}
