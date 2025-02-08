# DB migrations

## ER Diagram
```mermaid
erDiagram
    ACCOUNT {
        string id PK
        string name
        string email
    }

    ASSET {
        string id PK
        string symbol
        string name
        string market_code
    }

    TRANSACTION {
        string id PK
        string account_id FK
        string asset_id FK
        string transaction_type
        date transaction_date
        decimal quantity
        decimal price
        string currency_code
    }

    ASSET_PRICE {
        string asset_id PK, FK
        decimal price
        string currency_code
        timestamp trade_date PK
    }

    ASSET_BALANCE {
        string account_id PK, FK
        string asset_id PK, FK
        decimal quantity
    }

    CURRENCY_RATE {
        int id PK
        string base_currency
        string target_currency
        decimal exchange_rate
        timestamp trade_date
    }

    ACCOUNT ||--o{ TRANSACTION : "makes"
    ASSET ||--o{ TRANSACTION : "involves"
    ASSET ||--o{ ASSET_PRICE : "has"
    ACCOUNT ||--o{ ASSET_BALANCE : "holds"
    ASSET ||--o{ ASSET_BALANCE : "includes"

```

## How to install golang-migrate
```shell
brew install golang-migrate
```

## How to install mySQL
```shell
podman  pull mysql:9.2
podman run -p 3306:3306 --name atp-db -e MYSQL_ROOT_PASSWORD=Pass123  -d mysql:9.2
podman exec -it atp-db bash
```

Check if running
```shell
podman ps 
podman port atp-db
```

## Create new database
```sql
CREATE DATABASE atp_db;
```

## db migrations in the minikube kubernetes cluster
To run db migrations in the minikube kubernetes cluster, replace the .env values
```shell
make migrate_up
```


## References
- https://www.freecodecamp.org/news/database-migration-golang-migrate/