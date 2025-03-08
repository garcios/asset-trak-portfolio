WITH LatestPrices AS (
    (
        WITH ranked_prices AS (
            SELECT
                ap.asset_id,
                ap.price,
                ap.trade_date,
                ROW_NUMBER() OVER (PARTITION BY ap.asset_id ORDER BY ap.trade_date DESC) AS row_num
            FROM
                asset_price ap
            WHERE
                ap.trade_date <= '2024-03-05'
        )
        SELECT
            asset_id,
            price,
            trade_date AS latest_trade_date
        FROM
            ranked_prices
        WHERE
            row_num = 1
    )
)
SELECT
    t.asset_id,
    SUM(t.quantity) AS total_quantity,
    lp.price,
    t.trade_price_currency_code
FROM
    transaction t
        JOIN
    LatestPrices lp
    ON
        t.asset_id = lp.asset_id
WHERE
    t.transaction_date BETWEEN '2020-07-01' AND '2024-03-05'
  AND t.transaction_type IN ('BUY', 'SELL','SPLIT')
GROUP BY
    t.asset_id, lp.price, t.trade_price_currency_code;
