package service

type Config struct {
	Asset        Asset        `toml:"asset"`
	AssetPrice   AssetPrice   `toml:"asset_price"`
	CurrencyRate CurrencyRate `toml:"currency_rate"`
}

type AssetPrice struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
}

type Asset struct {
	Symbols  []string `toml:"symbols"`
	Path     string   `toml:"path"`
	SkipRows int      `toml:"skip_rows"`
}

type CurrencyRate struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
}
