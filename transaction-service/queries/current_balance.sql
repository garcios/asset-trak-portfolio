SELECT a.symbol, b.quantity FROM asset_balance b
JOIN asset a ON b.asset_id = a.id where quantity >0
ORDER BY a.symbol;