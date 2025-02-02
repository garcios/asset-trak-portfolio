# SQL Queries
This directory incorporates a collection of SQL queries, specifically designed for conducting data analysis. 
Additionally, these queries provide a testing framework before their integration into the main codebase.

## Get Summary Balance for each individual asset converted to AUD currency.
>Filename: current_balance_converted.sql
```sql
WITH bq AS (
    SELECT a.symbol, b.quantity, a.id AS asset_id FROM asset_balance b
                                                           JOIN asset a ON b.asset_id = a.id
    WHERE quantity >0
    ORDER BY a.symbol
)
SELECT  bq.symbol,
        bq.quantity,
        p.price,
        CASE
            WHEN p.currency_code = 'USD' THEN  bq.quantity* p.price*c.exchange_rate
        ELSE  bq.quantity*p.price
        END AS total,
        CASE
            WHEN p.currency_code = 'USD' THEN  c.exchange_rate
            ELSE  1
            END AS exchange_rate,
        c.trade_date AS currency_date,
        p.trade_date
FROM bq
LEFT JOIN (
    SELECT asset_id, price, trade_date, currency_code
    FROM (
             SELECT asset_id, price, trade_date, currency_code,
                    ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
             FROM asset_price) t
    WHERE t.rn = 1
) p ON p.asset_id = bq.asset_id
LEFT JOIN (
    SELECT exchange_rate, trade_date FROM currency_rate c
    WHERE base_currency='USD' AND target_currency='AUD'
      AND trade_date <= '2025-01-31' ORDER BY trade_date DESC LIMIT 1
) c ON 1 = 1
ORDER BY total DESC;
```

## Get total balance of all assets converted to AUD currency.
>Filename: current_balance_total.sql
```sql
WITH bq AS (
    SELECT a.symbol, b.quantity, a.id AS asset_id FROM asset_balance b
                                                           JOIN asset a ON b.asset_id = a.id
    WHERE quantity >0
    ORDER BY a.symbol
)
SELECT SUM(t.total) FROM (SELECT bq.symbol,
                                 bq.quantity,
                                 p.price,
                                 CASE
                                     WHEN p.currency_code = 'USD' THEN bq.quantity * p.price * c.exchange_rate
                                     ELSE bq.quantity * p.price
                                     END      AS total,
                                 CASE
                                     WHEN p.currency_code = 'USD' THEN c.exchange_rate
                                     ELSE 1
                                     END      AS exchange_rate,
                                 c.trade_date AS currency_date,
                                 p.trade_date
                          FROM bq
                                   LEFT JOIN (SELECT asset_id, price, trade_date, currency_code
                                              FROM (SELECT asset_id,
                                                           price,
                                                           trade_date,
                                                           currency_code,
                                                           ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
                                                    FROM asset_price) t
                                              WHERE t.rn = 1) p ON p.asset_id = bq.asset_id
                                   LEFT JOIN (SELECT exchange_rate, trade_date
                                              FROM currency_rate c
                                              WHERE base_currency = 'USD'
                                                AND target_currency = 'AUD'
                                                AND trade_date <= '2025-01-31'
                                              ORDER BY trade_date DESC LIMIT 1) c ON 1 = 1
                          ORDER BY total DESC) t;
```

