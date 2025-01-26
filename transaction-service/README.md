# Transaction service


## Processing excel file 
```shell
go get github.com/xuri/excelize/v2
```

## Install mysql driver package
```shell
go get -u github.com/go-sql-driver/mysql
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

## Run the asset ingestor
```shell
go run cmd/main.go -processor=assetIngestor
```

## Run the transaction ingestor
```shell
go run cmd/main.go -processor=transactionIngestor
```

## Run the transaction service gRPC
```shell
go run cmd/main.go 
```


## References
- https://www.kelche.co/blog/go/excel/
- 