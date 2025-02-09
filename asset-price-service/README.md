# Asset Price Service
The Asset Price Service is a core module in `Investment Portfolio Navigator` application, responsible for retrieving, 
processing, and managing real-time and historical asset prices. It ensures accurate valuation of portfolio holdings and 
enables data-driven decision-making.

__Key Features & Responsibilities__:
- Real-Time Price Updates: Continuously fetches live asset prices from financial data providers.
- Historical Price Data: Stores and maintains historical asset prices for performance analysis and backtesting.

By integrating the Asset Price Service, the `Investment Portfolio Navigator` application enhances accuracy, transparency, 
and efficiency, helping users track market movements and optimize investment strategies.

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

## References
- https://support.microsoft.com/en-au/office/stockhistory-function-1ac8b5b3-5f62-4d94-8ab8-7504ec7239a8