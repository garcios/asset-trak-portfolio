INSERT INTO portfolio_performance (account_id, trade_date, performance_5d, performance_1m, performance_6m, currency_code)
SELECT
    pv1.account_id,
    pv1.trade_date,
    ((pv1.total_value - pv5.total_value) / pv5.total_value) * 100 AS performance_5d,
    ((pv1.total_value - pv30.total_value) / pv30.total_value) * 100 AS performance_1m,
    ((pv1.total_value - pv180.total_value) / pv180.total_value) * 100 AS performance_6m,
    pv1.currency_code
FROM portfolio_value pv1
         LEFT JOIN portfolio_value pv5 ON pv5.account_id = pv1.account_id AND pv5.trade_date = DATE_SUB(pv1.trade_date, INTERVAL 5 DAY)
         LEFT JOIN portfolio_value pv30 ON pv30.account_id = pv1.account_id AND pv30.trade_date = DATE_SUB(pv1.trade_date, INTERVAL 1 MONTH)
         LEFT JOIN portfolio_value pv180 ON pv180.account_id = pv1.account_id AND pv180.trade_date = DATE_SUB(pv1.trade_date, INTERVAL 6 MONTH)
WHERE pv1.trade_date = CURDATE();
