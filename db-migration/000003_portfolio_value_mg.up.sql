CREATE TABLE portfolio_value (
  id INT AUTO_INCREMENT PRIMARY KEY,
  account_id VARCHAR(36) NOT NULL,
  trade_date DATE NOT NULL,
  total_value DECIMAL(20, 4) NOT NULL,
  currency_code VARCHAR(3) NOT NULL,
  FOREIGN KEY (account_id) REFERENCES account(id),
  UNIQUE (account_id, trade_date)
);

CREATE TABLE portfolio_performance (
  id INT AUTO_INCREMENT PRIMARY KEY,
  account_id VARCHAR(36) NOT NULL,
  trade_date DATE NOT NULL,
  performance_5d DECIMAL(10, 4),   -- Percentage change over the last 5 days
  performance_1m DECIMAL(10, 4),   -- Percentage change over the last 1 month
  performance_6m DECIMAL(10, 4),   -- Percentage change over the last 6 months
  performance_ytd DECIMAL(10, 4),   -- Percentage change over year to date
  performance_1y DECIMAL(10, 4),   -- Percentage change over the last 1 year
  performance_5y DECIMAL(10, 4),   -- Percentage change over the last 5 years
  FOREIGN KEY (account_id) REFERENCES account(id),
  UNIQUE (account_id, trade_date)
);
