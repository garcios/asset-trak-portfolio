SELECT a.symbol, b.quantity, p.price, b.quantity*p.price AS total, p.currency_code,  p.trade_date
FROM asset_balance b
         JOIN asset a ON b.asset_id = a.id
         LEFT JOIN (
    SELECT asset_id, price, trade_date, currency_code,
           ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
    FROM asset_price WHERE trade_date <= NOW()) p
                   ON p.asset_id = a.id AND p.rn = 1
WHERE b.quantity > 0
AND b.account_id = 'eb08df3c-958d-4ae8-b3ae-41ec04418786';

###
SELECT a.symbol, a.name, b.quantity, p.price, p.currency_code, a.market_code
FROM asset_balance b JOIN asset a ON b.asset_id = a.id
                     LEFT JOIN (
    SELECT asset_id, price, trade_date, currency_code,
           ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
    FROM asset_price WHERE trade_date <= NOW()) p
                               ON p.asset_id = a.id AND p.rn = 1
WHERE b.quantity > 0
  AND b.account_id = 'eb08df3c-958d-4ae8-b3ae-41ec04418786';