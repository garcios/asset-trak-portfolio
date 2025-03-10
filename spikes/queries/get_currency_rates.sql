SELECT base_currency, target_currency, exchange_rate, DATE(trade_date) FROM currency_rate
WHERE base_currency = 'USD'
  AND target_currency ='AUD'
  AND trade_date BETWEEN '2020-07-01' AND '2025-03-01';