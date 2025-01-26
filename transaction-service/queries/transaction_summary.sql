WITH transaction_summary AS (
    SELECT SUM(quantity) as qty, asset_id
    FROM transaction
    GROUP BY asset_id
    HAVING SUM(quantity) > 0
)
SELECT transaction_summary.qty, asset.symbol
FROM transaction_summary
  LEFT JOIN asset
  ON transaction_summary.asset_id = asset.id
ORDER BY asset.symbol;