WITH Balance_Quantity AS (
 SELECT a.symbol, b.quantity, a.id AS asset_id FROM asset_balance b
 JOIN asset a ON b.asset_id = a.id
  WHERE quantity >0
 ORDER BY a.symbol
 )
SELECT Balance_Quantity.asset_id, Balance_Quantity.symbol, Balance_Quantity.quantity
FROM Balance_Quantity





