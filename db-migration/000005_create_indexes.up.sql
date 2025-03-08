CREATE INDEX idx_transaction_asset_date_type
    ON transaction (asset_id, transaction_date, transaction_type);

CREATE INDEX idx_asset_price_asset_date
    ON asset_price (asset_id, trade_date DESC);

CREATE INDEX idx_transaction_type_filtered
    ON transaction (asset_id, transaction_date);


CREATE INDEX idx_asset_price_covering
    ON asset_price (asset_id, trade_date DESC, price);
