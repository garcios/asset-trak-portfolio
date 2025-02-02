# GraphQL

## Dependencies
```shell
go get -v github.com/garcios/asset-trak-portfolio/currency-service@main
```

## Run the GraphQL Server
```shell
go run server.go
```

## Open the graphQL playground
```shell
http://localhost:8080/
```

## Example Query
```graphql
query {
  getBalanceSummary(accountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786") {
    accountId
    totalValue {
      amount
      currencyCode
    }
    balanceItems {
      assetSymbol
      assetName
      price {
        amount
        currencyCode
      }
      quantity
      value {
        amount
        currencyCode
      }
      totalGain
      marketCode
    }
  }
}
```

## Example Response
```json
{
  "data": {
    "getBalanceSummary": {
      "accountId": "eb08df3c-958d-4ae8-b3ae-41ec04418786",
      "totalValue": {
        "amount": 321154.0082,
        "currencyCode": "AUD"
      },
      "balanceItems": [
        {
          "assetSymbol": "UBER",
          "assetName": "Uber Technologies Inc",
          "price": {
            "amount": 66.85,
            "currencyCode": "AUD"
          },
          "quantity": 52,
          "value": {
            "amount": 5596.682,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NYSE"
        },
        {
          "assetSymbol": "NVDA",
          "assetName": "NVIDIA Corp",
          "price": {
            "amount": 120.07,
            "currencyCode": "AUD"
          },
          "quantity": 45,
          "value": {
            "amount": 8699.0715,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "GOOG",
          "assetName": "Alphabet Inc - Ordinary Shares - Class C",
          "price": {
            "amount": 205.6,
            "currencyCode": "AUD"
          },
          "quantity": 5,
          "value": {
            "amount": 1655.0800000000002,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "WTC",
          "assetName": "Wisetech Global Ltd",
          "price": {
            "amount": 123.81,
            "currencyCode": "AUD"
          },
          "quantity": 30,
          "value": {
            "amount": 3714.3,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "ASX"
        },
        {
          "assetSymbol": "MSFT",
          "assetName": "Microsoft Corporation",
          "price": {
            "amount": 415.06,
            "currencyCode": "AUD"
          },
          "quantity": 49,
          "value": {
            "amount": 32744.0834,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "STW",
          "assetName": "Spdr S&P/Asx 200 Fund",
          "price": {
            "amount": 76.66,
            "currencyCode": "AUD"
          },
          "quantity": 347,
          "value": {
            "amount": 26601.02,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "ASX"
        },
        {
          "assetSymbol": "NDQ",
          "assetName": "Betashares Nasdaq 100 Etf",
          "price": {
            "amount": 51.6,
            "currencyCode": "AUD"
          },
          "quantity": 351,
          "value": {
            "amount": 18111.600000000002,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "ASX"
        },
        {
          "assetSymbol": "FTNT",
          "assetName": "Fortinet Inc",
          "price": {
            "amount": 100.88,
            "currencyCode": "AUD"
          },
          "quantity": 34,
          "value": {
            "amount": 5522.171200000001,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "AMAT",
          "assetName": "Applied Materials Inc.",
          "price": {
            "amount": 180.35,
            "currencyCode": "AUD"
          },
          "quantity": 15,
          "value": {
            "amount": 4355.4525,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "AVGO",
          "assetName": "Broadcom Inc",
          "price": {
            "amount": 221.27,
            "currencyCode": "AUD"
          },
          "quantity": 36,
          "value": {
            "amount": 12824.809200000002,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "AMD",
          "assetName": "Advanced Micro Devices Inc.",
          "price": {
            "amount": 115.95,
            "currencyCode": "AUD"
          },
          "quantity": 39,
          "value": {
            "amount": 7280.500500000001,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "MA",
          "assetName": "Mastercard Incorporated - Ordinary Shares - Class A",
          "price": {
            "amount": 555.43,
            "currencyCode": "AUD"
          },
          "quantity": 7,
          "value": {
            "amount": 6259.6961,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NYSE"
        },
        {
          "assetSymbol": "GOOGL",
          "assetName": "Alphabet Inc - Ordinary Shares - Class A",
          "price": {
            "amount": 204.02,
            "currencyCode": "AUD"
          },
          "quantity": 115,
          "value": {
            "amount": 37774.30300000001,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "AMZN",
          "assetName": "Amazon.com Inc.",
          "price": {
            "amount": 237.68,
            "currencyCode": "AUD"
          },
          "quantity": 116,
          "value": {
            "amount": 44389.1168,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NASDAQ"
        },
        {
          "assetSymbol": "TSM",
          "assetName": "Taiwan Semiconductor Manufacturing - ADR",
          "price": {
            "amount": 209.32,
            "currencyCode": "AUD"
          },
          "quantity": 35,
          "value": {
            "amount": 11795.182,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "NYSE"
        },
        {
          "assetSymbol": "IVV",
          "assetName": "Ishares S&P 500 Etf",
          "price": {
            "amount": 65.07,
            "currencyCode": "AUD"
          },
          "quantity": 1442,
          "value": {
            "amount": 93830.93999999999,
            "currencyCode": "AUD"
          },
          "totalGain": 0,
          "marketCode": "ASX"
        }
      ]
    }
  }
}
```


## References
- https://gqlgen.com/getting-started/
