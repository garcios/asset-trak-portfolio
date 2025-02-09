# Currency Service
The Currency Service Module is a critical component of `Investment Portfolio Navigator` application that ensures accurate 
handling of multi-currency assets, real-time exchange rates, and currency conversions. It enables users to manage global
investments efficiently by providing seamless currency-related functionalities.

Key Features & Responsibilities:
- Real-Time Exchange Rates: Fetches and updates live foreign exchange rates from trusted financial sources.
- Historical Currency Data: Stores past exchange rates for performance analysis and backtesting.
- Multi-Currency Portfolio Valuation: Converts portfolio holdings into a base currency for unified reporting.
- Currency Conversion Engine: Supports automatic and manual currency conversions for transactions and asset valuation.

By incorporating this module, the stock portfolio management application enhances accuracy, improves financial 
insights, and helps users optimize their international investments.

## Install toml parser
```shell
go get github.com/BurntSushi/toml
```

## Install my custom libraries
```shell
go get github.com/garcios/asset-trak-portfolio/lib@main 
```

Go-Micro V4
```shell
go get go-micro.dev/v4@latest
```

## Set env variables
```shell
export DBUSER=root
export DBPASS=Pass123
```


## References
- https://blog.apilayer.com/how-to-get-real-time-and-historical-exchange-rates-into-excel/#:~:text=To%20get%20historical%20exchange%20rates,%2C%2C0%2C1).