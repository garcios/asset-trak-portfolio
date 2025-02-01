package service

type Config struct {
	FileInfo FileInfo `toml:"file_info"`
}

type FileInfo struct {
	Path     string `toml:"path"`
	SkipRows int    `toml:"skip_rows"`
}
