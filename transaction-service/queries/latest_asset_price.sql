SELECT price, currency_code
FROM asset_price
WHERE asset_id = 'AMZN'
  AND timestamp <= '2025-01-02'
ORDER BY timestamp DESC
    LIMIT 1;