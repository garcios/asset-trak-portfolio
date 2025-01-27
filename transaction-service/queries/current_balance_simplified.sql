SELECT a.symbol, b.quantity, p.price, b.quantity*p.price AS total, p.currency_code,  p.trade_date
FROM asset_balance b
         JOIN asset a ON b.asset_id = a.id
         LEFT JOIN (
    SELECT asset_id, price, trade_date, currency_code,
           ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
    FROM asset_price WHERE trade_date <= '2025-01-15') p
                   ON p.asset_id = a.id AND p.rn = 1
WHERE b.quantity > 0
ORDER BY a.symbol;