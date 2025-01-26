SELECT exchange_rate
FROM currency_rate
WHERE base_currency = 'USD'
  AND target_currency = 'AUD'
  AND updated_at <= '2025-01-02'
ORDER BY updated_at DESC
    LIMIT 1;