SELECT asset_id, price, trade_date, currency_code
FROM (
         SELECT asset_id, price, trade_date, currency_code,
                ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
         FROM asset_price WHERE trade_date <= '2025-01-02') t
WHERE t.rn = 1