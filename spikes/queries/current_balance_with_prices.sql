WITH bq AS (
    SELECT a.symbol, b.quantity, a.id AS asset_id FROM asset_balance b
                                                           JOIN asset a ON b.asset_id = a.id
    WHERE quantity >0
    ORDER BY a.symbol
)
SELECT  bq.symbol, bq.quantity, p.price, bq.quantity*p.price AS total, p.currency_code,  p.trade_date
FROM bq
LEFT JOIN (
    SELECT asset_id, price, trade_date, currency_code
    FROM (
             SELECT asset_id, price, trade_date, currency_code,
                    ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
             FROM asset_price) t
    WHERE t.rn = 1
) p ON p.asset_id = bq.asset_id







