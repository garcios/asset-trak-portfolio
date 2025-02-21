package service

type Config struct {
	Asset    Asset    `toml:"asset"`
	FileInfo FileInfo `toml:"file_info"`
}

type FileInfo struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
}

type Asset struct {
	Symbols []string `toml:"symbols"`
}
