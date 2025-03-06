package service

type Config struct {
	Asset        Asset        `toml:"asset"`
	AssetPrice   AssetPrice   `toml:"asset_price"`
	CurrencyRate CurrencyRate `toml:"currency_rate"`
}

type AssetPrice struct {
	DirPath string `toml:"dir_path"`
}

type Asset struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
}

type CurrencyRate struct {
	DirPath string `toml:"dir_path"`
}
