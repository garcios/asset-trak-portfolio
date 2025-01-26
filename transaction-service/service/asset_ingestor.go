package service

import (
	"github.com/garcios/asset-trak-portfolio/transactions-service/db"
	"github.com/garcios/asset-trak-portfolio/transactions-service/model"
	"log"

	"github.com/google/uuid"
)

type IAssetManager interface {
	AddAsset(rec *model.Asset) error
	AssetExists(symbol string, marketCode string) (bool, error)
	FindAssetBySymbol(symbol string) (*model.Asset, error)
}

// verify interface compliance
var _ IAssetManager = &db.AssetRepository{}

type AssetIngestor struct {
	AssetManager IAssetManager
}

func NewAssetIngestor(am IAssetManager) *AssetIngestor {
	return &AssetIngestor{
		AssetManager: am,
	}
}

func (ingestor *AssetIngestor) ProcessAssets(filePath string, tabName string, skipRows int) error {
	log.Println("Processing assets...")
	rows, err := getRows(filePath, tabName)
	if err != nil {
		return err
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < skipRows {
			rowCount++
			continue
		}

		if len(row) == 0 || row[0] == "Total" {
			break
		}

		asset, err := mapColumnsToAsset(row)
		if err != nil {
			return err
		}

		isFound, err := ingestor.AssetManager.AssetExists(asset.Symbol, asset.MarketCode)
		if err != nil {
			return err
		}

		if isFound {
			continue
		}

		err = ingestor.AssetManager.AddAsset(asset)
		if err != nil {
			return err
		}

	}

	return nil
}

func mapColumnsToAsset(row []string) (*model.Asset, error) {
	id := uuid.New()
	rec := &model.Asset{ID: id.String(), Symbol: row[0], Name: row[2], MarketCode: row[1]}

	return rec, nil

}
