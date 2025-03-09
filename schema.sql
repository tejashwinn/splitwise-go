CREATE TABLE IF NOT EXISTS SW_USR (
    OBJECT_ID BIGSERIAL PRIMARY KEY,
    USR_NAME VARCHAR (50) UNIQUE NOT NULL,
    USR_PASSWORD VARCHAR (50) NOT NULL,
    USR_EMAIL VARCHAR (255) UNIQUE NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT TIMESTAMP
);
CREATE TABLE SW_TXN_CUR (
    OBJECT_ID BIGSERIAL PRIMARY KEY,
    CUR_CODE CHAR(3) UNIQUE NOT NULL,
    CUR_NAME TEXT UNIQUE NOT NULL,
    CUR_SYMBOL TEXT,
    CUR_EX_RATE NUMERIC(18, 6) NOT NULL CHECK (CUR_EX_RATE > 0),
    CUR_BASE_YN BOOLEAN NOT NULL DEFAULT FALSE,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT TIMESTAMP
);
INSERT INTO SW_TXN_CUR (
        CUR_CODE,
        CUR_NAME,
        CUR_SYMBOL,
        CUR_EX_RATE,
        CUR_BASE_YN
    )
VALUES (
        'USD',
        'United States Dollar',
        '$',
        1.000000,
        TRUE
    ),
    ('EUR', 'Euro', '€', 0.920000, FALSE),
    (
        'GBP',
        'British Pound Sterling',
        '£',
        0.780000,
        FALSE
    ),
    ('JPY', 'Japanese Yen', '¥', 150.250000, FALSE),
    ('INR', 'Indian Rupee', '₹', 83.500000, FALSE),
    (
        'AUD',
        'Australian Dollar',
        'A$',
        1.520000,
        FALSE
    ),
    ('CAD', 'Canadian Dollar', 'C$', 1.350000, FALSE),
    ('CHF', 'Swiss Franc', 'CHF', 0.890000, FALSE),
    ('CNY', 'Chinese Yuan', '¥', 7.150000, FALSE),
    ('SGD', 'Singapore Dollar', 'S$', 1.340000, FALSE);
