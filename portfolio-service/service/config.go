package service

type Config struct {
	Account   Account   `toml:"account"`
	Trades    Trades    `toml:"trades"`
	Dividends Dividends `toml:"dividends"`
}

type Account struct {
	ID string `toml:"id"`
}

type Trades struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
	TabName  string `toml:"tab_name"`
}

type Dividends struct {
	Path             string `toml:"path"`
	TabNameDomestic  string `toml:"tab_name_domestic"`
	SkipRowsDomestic int    `toml:"skip_rows_domestic"`
	TabNameForeign   string `toml:"tab_name_foreign"`
	SkipRowsForeign  int    `toml:"skip_rows_foreign"`
}
