CREATE SCHEMA sdetozon;

CREATE TABLE sdetozon.mock_scenarios (
    test_id     VARCHAR(255) PRIMARY KEY,
    status_code INTEGER NOT NULL
);

CREATE TABLE sdetozon.exchange_rates (
    rate_id              VARCHAR(255) NOT NULL,
    num_code             VARCHAR(3) NOT NULL,
    char_code            VARCHAR(3) NOT NULL,
    nominal              INTEGER NOT NULL,
    value_name           VARCHAR(255) NOT NULL,
    value                VARCHAR(255) NOT NULL,
    vunit_rate           VARCHAR(255) NOT NULL,

    test_id VARCHAR(255) PRIMARY KEY REFERENCES sdetozon.mock_scenarios(test_id) ON DELETE CASCADE
);
