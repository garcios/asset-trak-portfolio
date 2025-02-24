package service

import (
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"github.com/patrickmn/go-cache"
	"time"
)

type allCache struct {
	assets *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

// getAsset retrieves an asset by its symbol. If the asset is not found in the cache, it is retrieved from the database.
func (ingestor *TransactionIngestor) getAsset(assetSymbol string) (*model.Asset, error) {
	if assetFromCache, ok := ingestor.cache.assets.Get(assetSymbol); ok {
		return assetFromCache.(*model.Asset), nil
	}

	asset, err := ingestor.AssetManager.FindAssetBySymbol(assetSymbol)
	if err != nil {
		return nil, err
	}

	ingestor.cache.assets.Set(assetSymbol, asset, cache.DefaultExpiration)

	return asset, nil
}
