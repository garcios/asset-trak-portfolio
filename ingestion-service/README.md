# Ingestion Service
The Ingestion Service is a core module in `Investment Portfolio Navigator` application, responsible for ingestion of 
market data such as assets, asset prices, and currency rates.

__Key Features & Responsibilities__:
- Real-Time Price Updates: Continuously fetches live asset prices from financial data providers.
- Historical Price Data: Stores and maintains historical asset prices for performance analysis and backtesting.

## Install toml parser
```shell
go get github.com/BurntSushi/toml
```

## Install my custom libraries
```shell
go get github.com/garcios/asset-trak-portfolio/lib@main 
```

## Set env variables
```shell
export DBUSER=root
export DBPASS=Pass123
```

## Market data
- US stocks prices and FX rates were downloaded from https://www.alphavantage.co
- AU stocks were downloaded from https://www.marketindex.com.au


## References
- <https://support.microsoft.com/en-au/office/stockhistory-function-1ac8b5b3-5f62-4d94-8ab8-7504ec7239a8>
- <https://www.marketindex.com.au/>
- <https://www.alphavantage.co/>
