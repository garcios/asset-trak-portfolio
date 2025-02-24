CREATE TABLE account (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(50)  NOT NULL,
  email VARCHAR(50)  NOT NULL
);

CREATE TABLE asset (
   id  VARCHAR(36) PRIMARY KEY,
   symbol VARCHAR(50) NOT NULL,
   name VARCHAR(255) NOT NULL,
   market_code  VARCHAR(20) NOT NULL
);

CREATE TABLE transaction (
  id VARCHAR(36) PRIMARY KEY,
  account_id VARCHAR(36) NOT NULL,
  asset_id VARCHAR(36),  -- Optional: Associated with an asset (e.g., for dividends, buy/sell)
  transaction_type VARCHAR(10)  NOT NULL, -- E.g., "BUY", "SELL", "DIVIDEND", "INTEREST"
  transaction_date DATE  NOT NULL, -- trade date or date of dividend payout
  quantity DECIMAL(15, 4), -- Amount of shares or units (optional for non-asset transactions)
  trade_price DECIMAL(15, 4), -- Optional: purchase or sell price
  trade_price_currency_code VARCHAR(3), -- ISO currency code of the transaction price
  brokerage_fee DECIMAL(15, 4), -- Fee charged for processing the transaction
  fee_currency_code VARCHAR(3), -- ISO currency code for the brokerage fee
  amount_cash DECIMAL(15, 4), -- Optional: Cash amount (e.g., for dividends)
  amount_currency_code VARCHAR(3),  -- ISO currency code for cash amount
  exchange_rate DECIMAL(15, 8) NOT NULL, -- Conversion rate for currency
  withheld_tax_amount DECIMAL(15, 4), -- Optional: Tax withheld (e.g., for dividends)
  withheld_tax_currency_code VARCHAR(3),  -- ISO currency code for tax withheld
  FOREIGN KEY (account_id) REFERENCES account(id),
  FOREIGN KEY (asset_id) REFERENCES asset(id)
);

CREATE TABLE asset_price (
   asset_id VARCHAR(36)  NOT NULL,
   price DECIMAL(15, 2) NOT NULL,
   currency_code   VARCHAR(3)  NOT NULL,
   trade_date TIMESTAMP NOT NULL,
   PRIMARY KEY (asset_id, trade_date),
   FOREIGN KEY (asset_id) REFERENCES asset(id)
);

CREATE TABLE asset_balance (
   account_id VARCHAR(36),
   asset_id VARCHAR(36)  NOT NULL,
   quantity DECIMAL(15, 4) NOT NULL,
   PRIMARY KEY (account_id, asset_id),
   FOREIGN KEY (account_id) REFERENCES account(id),
   FOREIGN KEY (asset_id) REFERENCES asset(id)
);

CREATE TABLE currency_rate (
 id INT AUTO_INCREMENT PRIMARY KEY,
 base_currency VARCHAR(10) NOT NULL,
 target_currency VARCHAR(10) NOT NULL,
 exchange_rate DECIMAL(15, 4) NOT NULL,
 trade_date TIMESTAMP  NOT NULL
);