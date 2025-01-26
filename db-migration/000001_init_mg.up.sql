CREATE TABLE account (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(50)  NOT NULL,
  email VARCHAR(50)  NOT NULL
);

CREATE TABLE transaction (
  id VARCHAR(36) PRIMARY KEY,
  account_id VARCHAR(36) NOT NULL,
  asset_symbol VARCHAR(10)  NOT NULL,
  transaction_type VARCHAR(10)  NOT NULL,
  transaction_date DATE  NOT NULL,
  quantity INT  NOT NULL,
  price DECIMAL(10, 2)  NOT NULL,
  currency_code   VARCHAR(3)  NOT NULL,
  FOREIGN KEY (account_id) REFERENCES account(id)
);

CREATE TABLE asset (
    id  VARCHAR(36) PRIMARY KEY,
    symbol VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    market_code  VARCHAR(20) NOT NULL
);

CREATE TABLE asset_prices (
   asset_id VARCHAR(36)  NOT NULL,
   price DECIMAL(15, 2) NOT NULL,
   currency_code   VARCHAR(3)  NOT NULL,
   timestamp TIMESTAMP NOT NULL,
   PRIMARY KEY (asset_id, timestamp),
   FOREIGN KEY (asset_id) REFERENCES asset(id)
);

CREATE TABLE asset_balance (
   account_id VARCHAR(36),
   asset_id VARCHAR(36)  NOT NULL,
   balance DECIMAL(15, 2) NOT NULL,
   currency_code   VARCHAR(3)  NOT NULL,
   PRIMARY KEY (account_id, asset_id),
   FOREIGN KEY (account_id) REFERENCES account(id),
   FOREIGN KEY (asset_id) REFERENCES asset(id)
);

