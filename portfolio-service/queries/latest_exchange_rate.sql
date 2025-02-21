SELECT exchange_rate
FROM currency_rate
WHERE base_currency = 'USD'
  AND target_currency = 'AUD'
  AND trade_date <= '2025-01-02'
ORDER BY trade_date DESC
    LIMIT 1;