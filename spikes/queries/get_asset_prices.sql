SELECT asset_id, price, currency_code, DATE(trade_date) FROM asset_price
WHERE asset_id = '0687032c-52f0-4700-8f5e-5ba626ff56f9'
AND trade_date BETWEEN '2020-07-01' AND '2025-03-01';