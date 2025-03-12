module github.com/garcios/asset-trak-portfolio/ingestion-service

go 1.23

toolchain go1.23.6

replace github.com/garcios/asset-trak-portfolio/lib => /Users/oscargarcia/workspace/asset-trak-portfolio/lib

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/google/uuid v1.6.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/garcios/asset-trak-portfolio/lib v0.0.0-20250222045010-be1f0e8df8a9
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)
