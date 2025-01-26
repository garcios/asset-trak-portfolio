INSERT INTO asset (id, symbol, name, market_code)
VALUES
    ('4ecc5b0a-15ff-4c7b-83cf-2ef6dc7f9a5','AMZN', 'Amazon.com, Inc', 'NASDAQ'),
    ('e66a8281-5acd-45d7-97d4-aac1a1b9b70a','GOOGL', 'Alphabet Inc', 'NASDAQ'),
    ('acc24285-fdf9-461d-9d6c-a7ad74e01d5d','MSFT', 'Microsoft Corporation', 'NASDAQ'),
    ('b742c0a6-f66b-4ca2-8ca0-21b15b6d7895','AVGO', 'Broadcom Inc', 'NASDAQ'),
    ('c5b2ea8e-e27f-4bae-b40f-d30d9d9cb953','TSM', 'Taiwan Semiconductor Mfg. Co.','NYSE'),
    ('704441f9-7cd4-472c-a043-79537b274ae6','NVIDIA', 'NVDIA Corp', 'NASDAQ'),
    ('6caa9de2-a922-4d10-b975-7f84fe5d15de','AMD', 'Advanced Micro Devices Inc', 'NASDAQ'),
    ('1bbe0bb5-ea73-45d3-b0c9-c4f2c4840492','MA', 'Mastercard Inc','NYSE'),
    ('f7d8e0ce-dcae-4807-ad28-d87a09ebb0f2','UBER', 'Uber Technologies Inc','NYSE'),
    ('83eed873-579e-4f2f-b8f8-cbc3d77804f2','FTNT', 'Fortinet Inc', 'NASDAQ'),
    ('116b5c70-d25e-4741-8b04-281a4fabb18f','AMAT', 'Applied Materials Inc', 'NASDAQ')
;

INSERT INTO account (id, name, email)
VALUES
    ('eb08df3c-958d-4ae8-b3ae-41ec04418786','Harry Potter', 'harry@wizard.com'),
    ('2e633070-d368-435c-a01b-b382c9d94681','Hermione Granger', 'hermione@wizard.com'),
    ('6aee1937-336e-41fe-9f50-01788301caf4','Ron Weasley', 'ron@wizard.com')
;

COMMIT;